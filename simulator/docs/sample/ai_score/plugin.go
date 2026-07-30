package ai_score

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type AIScore struct {
	handle              framework.Handle
	httpClient          *http.Client
	modelURL            string
	calibrateURL        string
	confidenceThreshold float64

	// filterMinScore — жёсткий порог для Filter(): ноды с RF-баллом ниже
	// этого значения отбраковываются ДАЖЕ ЕСЛИ физически вмещают под
	// (в отличие от дефолтного NodeResourcesFit, который смотрит только на
	// сырые ресурсы). По умолчанию ВЫКЛЮЧЕНО (-1 = Filter всегда пропускает)
	// — намеренно консервативный дефолт, чтобы включение Filter не сломало
	// молча уже работающие бенчмарки Разделов 6.6-6.11; включается явно
	// через ML_FILTER_MIN_SCORE.
	filterMinScore int64
	// permitWaitDuration — на сколько Permit() приостанавливает Bind при
	// низкой calibrated confidence (даёт framework шанс пересмотреть решение
	// или дождаться более свежих данных о занятости кластера).
	permitWaitDuration time.Duration
	// maxPermitRetries — ЖЁСТКИЙ ПРЕДОХРАНИТЕЛЬ от бесконечного цикла
	// Reserve→Permit(Wait)→timeout→Unreserve→retry. CycleState создаётся
	// framework'ом заново на КАЖДУЮ попытку планирования пода и не переживает
	// retry — поэтому счётчик попыток здесь хранится ВНЕ CycleState, в самом
	// плагине, по UID пода. Обнаружено эмпирически: под со структурно низкой
	// (не меняющейся между попытками) confidence зацикливался на несколько
	// минут без этого предохранителя, прежде чем был остановлен вручную.
	maxPermitRetries int
	// permitAttempts — ключ НЕ UID пода, а namespace/name (см. podKey()):
	// эмпирически обнаружено, что этот облегчённый KWOK-based симулятор не
	// гарантирует стабильность p.UID между повторными выборками пода из
	// informer cache между попытками планирования — счётчик по UID никогда
	// не растущий выше 1 был первым, неудачным вариантом этого предохранителя.
	permitAttempts map[string]int
	permitMu       sync.Mutex

	// localityURL — адрес sidecar-сервиса, отдающего локацию блоков Alluxio
	// по пути к файлу (сам sidecar shell'ит `alluxio fs stat`, поскольку
	// готового REST-эндпоинта для block-locations в этой сборке Alluxio не
	// обнаружено — /api/v1/overview возвращает 404, /api/v1/master/ отдаёт
	// веб-фронтенд, а не JSON). См. также design-документ связки Модуль 1 —
	// Модуль 2, Раздел 9.
	localityURL string
	// localityBonus — фиксированная прибавка к итоговому score, если данные
	// пода (путь, взятый из аннотации DataPathAnnotationKey) уже кэшированы
	// на оцениваемой ноде. НЕ основано на калибровке — подобрано "на глаз"
	// как разумная отправная точка; требует последующей калибровки на
	// реальном распределении score, аналогично confidenceThreshold выше.
	localityBonus int64
	// localityAnnotationKey — ключ аннотации пода, где ожидается путь к
	// данным Alluxio (например, "alluxio.io/data-path"). ВАЖНО: это
	// синтетическая связка для целей эксперимента — реальные workload'ы,
	// использованные в основных экспериментах AIScore (Разделы 6.6-6.11
	// мастер-отчёта), такой аннотации не содержат; она должна быть явно
	// проставлена в тестовых манифестах.
	localityAnnotationKey string
	// localityHTTPClient — отдельный клиент с увеличенным таймаутом:
	// alluxio fs stat через sidecar эмпирически занимает ~2.2с (JVM CLI
	// стартует заново на каждый вызов), что превышает 2с таймаут pl.httpClient,
	// настроенный под быстрый ML-инференс. Общий клиент не трогаем, чтобы не
	// изменить поведение уже проверенных /predict и /calibrate_confidence.
	localityHTTPClient *http.Client

	replicationURL        string
	replicationMinNodes   int
	pathNodes             map[string]map[string]bool
	pathNodesMu           sync.Mutex
	replicationHTTPClient *http.Client
}


// podKey — стабильный (в отличие от p.UID на этом стенде) идентификатор
// пода для счётчика попыток Permit.
func podKey(p *v1.Pod) string {
	return p.Namespace + "/" + p.Name
}

var _ framework.ScorePlugin = &AIScore{}
var _ framework.FilterPlugin = &AIScore{}
var _ framework.ReservePlugin = &AIScore{}
var _ framework.PermitPlugin = &AIScore{}

const Name = "AIScore"

// confidenceStateKey — ключ CycleState, которым NormalizeScore передаёт
// посчитанную calibrated confidence дальше по конвейеру framework, в Permit.
// CycleState существует ровно на время одного цикла Schedule для одного
// пода, поэтому это безопасный, стандартный способ передать данные между
// точками расширения без глобального состояния.
const confidenceStateKey = "AIScore/confidence"

func (pl *AIScore) Name() string {
	return Name
}
type replicationRequest struct {
	Path     string `json:"path"`
	Replicas int    `json:"replicas"`
}

type replicationResponse struct {
	Path     string `json:"path"`
	Replicas int    `json:"replicas"`
	Success  bool   `json:"success"`
	Message  string `json:"message"`
}

type localityResponse struct {
	Path                string   `json:"path"`
	Found               bool     `json:"found"`
	Hosts               []string `json:"hosts"`
	Tiers               []string `json:"tiers"`
	InAlluxioPercentage int      `json:"in_alluxio_percentage"`
}
type predictRequest struct {
	NodeFreeCPU  int64 `json:"node_free_cpu"`
	NodeFreeMem  int64 `json:"node_free_mem"`
	NodeAllocCPU int64 `json:"node_alloc_cpu"`
	NodeAllocMem int64 `json:"node_alloc_mem"`
	PodReqCPU    int64 `json:"pod_req_cpu"`
	PodReqMem    int64 `json:"pod_req_mem"`
	PodPriority  int32 `json:"pod_priority"`
}

type predictResponse struct {
	Score int64 `json:"score"`
}

type calibrateRequest struct {
	Margin float64 `json:"margin"`
}

type calibrateResponse struct {
	Confidence float64 `json:"confidence"`
	Calibrated bool    `json:"calibrated"`
}

// confidenceStateData — обёртка над calibrateResponse, реализующая
// framework.StateData (единственный метод — Clone()), чтобы её можно было
// положить в CycleState.Write() и прочитать в Permit() тем же ключом.
type confidenceStateData struct {
	confidence float64
	calibrated bool
}

func (c *confidenceStateData) Clone() framework.StateData {
	return &confidenceStateData{confidence: c.confidence, calibrated: c.calibrated}
}

// heuristicScore — та же логика, что AIScoreHeuristic (max free CPU) и
// heuristic_score() в run_aiscore_over_scenarios.py — используется как
// fallback, когда calibrated confidence ниже порога (Раздел 6.9 мастер-отчёта).
// Балл нормализован в [0, framework.MaxNodeScore] относительно allocatable
// CPU ноды, чтобы быть сопоставимым по шкале с RF-баллом (который тоже
// клэмпится в [0,100] в main.py).
func heuristicScore(nodeInfo *framework.NodeInfo) int64 {
	allocatable := nodeInfo.Allocatable
	requested := nodeInfo.Requested
	freeCPU := allocatable.MilliCPU - requested.MilliCPU
	if freeCPU < 0 {
		freeCPU = 0
	}
	if allocatable.MilliCPU <= 0 {
		return 50
	}
	score := int64(float64(freeCPU) / float64(allocatable.MilliCPU) * float64(framework.MaxNodeScore))
	if score > framework.MaxNodeScore {
		score = framework.MaxNodeScore
	}
	if score < framework.MinNodeScore {
		score = framework.MinNodeScore
	}
	return score
}

// extractFeatures — переводит сырые ресурсы Kubernetes (millicores, байты) в
// конвенцию Alibaba Cluster Trace, на которой обучена модель: CPU в "100
// единиц = 1 ядро" (НЕ millicores!), память — % от allocatable ноды (не
// абсолютные байты). БЕЗ этого пересчёта модель получает вход на 2-3
// порядка вне диапазона обучения — деревья RF почти всегда попадают в один
// и тот же крайний лист, давая почти-константный балл (эмпирически: 31 или
// 50 почти на любом входе). Тот же баг уже был найден и исправлен в offline
// run_aiscore_over_scenarios.py (--rescale-to-trace-units, Раздел 6.9.2-6.9.3
// мастер-отчёта) — здесь применяется тот же пересчёт, ранее пропущенный в
// самом плагине, из-за чего ВСЯ live-серия экспериментов (Разделы 6.6,
// 6.10, 6.11) была получена на фактически неработающей модели.
func extractFeatures(pod *v1.Pod, nodeInfo *framework.NodeInfo) predictRequest {
	allocatable := nodeInfo.Allocatable
	requested := nodeInfo.Requested

	freeCPUMilli := allocatable.MilliCPU - requested.MilliCPU
	freeMemBytes := allocatable.Memory - requested.Memory
	if freeCPUMilli < 0 {
		freeCPUMilli = 0
	}
	if freeMemBytes < 0 {
		freeMemBytes = 0
	}

	var podReqCPUMilli, podReqMemBytes int64
	for _, c := range pod.Spec.Containers {
		podReqCPUMilli += c.Resources.Requests.Cpu().MilliValue()
		podReqMemBytes += c.Resources.Requests.Memory().Value()
	}

	var priority int32
	if pod.Spec.Priority != nil {
		priority = *pod.Spec.Priority
	}

	// CPU: millicores -> "100 единиц = 1 ядро" (1000m = 1 ядро => /10)
	nodeFreeCPUTrace := freeCPUMilli / 10
	nodeAllocCPUTrace := allocatable.MilliCPU / 10
	podReqCPUTrace := podReqCPUMilli / 10

	// Память: приближение как % от allocatable ноды (точный эталон
	// нормализации Alibaba Trace недоступен — см. threats to validity
	// Раздела 6.9.4; это тот же proxy, что и в offline-скрипте).
	var nodeFreeMemPct, podReqMemPct int64
	if allocatable.Memory > 0 {
		nodeFreeMemPct = int64(float64(freeMemBytes) / float64(allocatable.Memory) * 100)
		podReqMemPct = int64(float64(podReqMemBytes) / float64(allocatable.Memory) * 100)
	}

	return predictRequest{
		NodeFreeCPU:  nodeFreeCPUTrace,
		NodeFreeMem:  nodeFreeMemPct,
		NodeAllocCPU: nodeAllocCPUTrace,
		NodeAllocMem: 100, // allocatable ноды = 100% по определению в этой шкале
		PodReqCPU:    podReqCPUTrace,
		PodReqMem:    podReqMemPct,
		PodPriority:  priority,
	}
}

func (pl *AIScore) callMLService(ctx context.Context, req predictRequest) (int64, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, pl.modelURL, bytes.NewReader(data))
	if err != nil {
		return 0, fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := pl.httpClient.Do(httpReq)
	if err != nil {
		return 0, fmt.Errorf("http call: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("ml-server returned status %d", resp.StatusCode)
	}

	var pr predictResponse
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return 0, fmt.Errorf("decode response: %w", err)
	}
	return pr.Score, nil
}

// callLocalityService спрашивает sidecar, на каких хостах и в каком тире
// уже закэшированы данные по указанному пути Alluxio. Возвращает
// (nil, nil) если localityURL не сконфигурирован — вызывающий код должен
// трактовать это как "бонус не применяется", а не как ошибку.
func (pl *AIScore) callLocalityService(ctx context.Context, path string) (*localityResponse, error) {
	if pl.localityURL == "" || path == "" {
		return nil, nil
	}

	reqURL := fmt.Sprintf("%s?path=%s", pl.localityURL, path)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("build locality request: %w", err)
	}

	resp, err := pl.localityHTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("locality http call: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("locality sidecar returned status %d", resp.StatusCode)
	}

	var lr localityResponse
	if err := json.NewDecoder(resp.Body).Decode(&lr); err != nil {
		return nil, fmt.Errorf("decode locality response: %w", err)
	}
	return &lr, nil
}

// nodeMatchesLocality проверяет, входит ли NodeInternalIP ноды в список
// хостов, где уже закэшированы данные. Ранее в этом KWOK-стенде ноды,
// созданные через kwokctl, не имели NodeInternalIP вообще (только
// Hostname=localhost, одинаковый у всех) — это временно решалось
// упрощением "бонус для любой ноды". По прецеденту Модуля 1 (создание
// гетерогенных нод через прямой PUT .../status с явно заданными
// allocatable/capacity) ноды для этого эксперимента создаются с явно
// заданным status.addresses, включая NodeInternalIP — что делает реальное
// сопоставление по IP снова содержательным и информативным.
func nodeMatchesLocality(nodeInfo *framework.NodeInfo, lr *localityResponse) bool {
	if lr == nil || !lr.Found {
		return false
	}
	node := nodeInfo.Node()
	if node == nil {
		return false
	}
	for _, addr := range node.Status.Addresses {
		if addr.Type != v1.NodeInternalIP {
			continue
		}
		for _, host := range lr.Hosts {
			if addr.Address == host {
				return true
			}
		}
	}
	return false
}

// callReplicationService просит sidecar установить минимальное число реплик
// для указанного пути. Ошибка не должна ронять Reserve() — репликация это
// оптимизация, а не обязательное условие для успешного размещения пода.
func (pl *AIScore) callReplicationService(ctx context.Context, path string, replicas int) error {
	if pl.replicationURL == "" || path == "" {
		return nil
	}

	reqBody := replicationRequest{Path: path, Replicas: replicas}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("marshal replication request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, pl.replicationURL, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("build replication request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := pl.replicationHTTPClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("replication http call: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("replication sidecar returned status %d", resp.StatusCode)
	}

	var rr replicationResponse
	if err := json.NewDecoder(resp.Body).Decode(&rr); err != nil {
		return fmt.Errorf("decode replication response: %w", err)
	}
	if !rr.Success {
		return fmt.Errorf("replication sidecar reported failure: %s", rr.Message)
	}
	return nil
}

// recordPodPlacement добавляет ноду в множество нод, на которые размещены
// поды с данным data-path, и возвращает текущее число РАЗЛИЧНЫХ нод.
func (pl *AIScore) recordPodPlacement(path, nodeName string) int {
	pl.pathNodesMu.Lock()
	defer pl.pathNodesMu.Unlock()

	if pl.pathNodes == nil {
		pl.pathNodes = make(map[string]map[string]bool)
	}
	if pl.pathNodes[path] == nil {
		pl.pathNodes[path] = make(map[string]bool)
	}
	pl.pathNodes[path][nodeName] = true
	return len(pl.pathNodes[path])
}

// callCalibrateService отправляет margin (разрыв между лучшим и вторым по
// баллу кандидатом для одного пода) и получает откалиброванную уверенность.
// Ошибка здесь не должна ронять Score-раунд целиком — вызывающий код
// (NormalizeScore) при ошибке пропускает шаг нормализации целиком и
// оставляет исходные RF-баллы как есть (эквивалент поведения "confidence
// не измерялась", а не присвоение конкретного числового confidence).
func (pl *AIScore) callCalibrateService(ctx context.Context, margin float64) (calibrateResponse, error) {
	req := calibrateRequest{Margin: margin}
	data, err := json.Marshal(req)
	if err != nil {
		return calibrateResponse{}, fmt.Errorf("marshal calibrate request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, pl.calibrateURL, bytes.NewReader(data))
	if err != nil {
		return calibrateResponse{}, fmt.Errorf("build calibrate request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := pl.httpClient.Do(httpReq)
	if err != nil {
		return calibrateResponse{}, fmt.Errorf("http calibrate call: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return calibrateResponse{}, fmt.Errorf("ml-server calibrate returned status %d", resp.StatusCode)
	}

	var cr calibrateResponse
	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		return calibrateResponse{}, fmt.Errorf("decode calibrate response: %w", err)
	}
	return cr, nil
}

// Filter — точка расширения #1 из 4 (была не реализована). Отбраковывает
// ноду ДО стадии Score, если её RF-балл ниже жёсткого порога, даже когда
// нода технически вмещает под по сырым ресурсам (это отличает Filter от
// уже существующего дефолтного NodeResourcesFit). ВЫКЛЮЧЕНО по умолчанию
// (filterMinScore=-1) — намеренно консервативно, чтобы не сломать молча уже
// работающие бенчмарки Разделов 6.6-6.11 при простом добавлении секции
// "filter" в scheduler.yaml.
func (pl *AIScore) Filter(ctx context.Context, state *framework.CycleState,
	pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {

	if pl.filterMinScore < 0 {
		return nil // Filter выключен — пропускаем всех кандидатов
	}

	features := extractFeatures(pod, nodeInfo)
	score, err := pl.callMLService(ctx, features)
	if err != nil {
		klog.Warningf("AIScore Filter fallback: node=%s pod=%s err=%v — пропускаем ноду, "+
			"не блокируем на ошибке ML-сервиса", nodeInfo.Node().Name, pod.Name, err)
		return nil
	}

	if score < pl.filterMinScore {
		return framework.NewStatus(framework.Unschedulable,
			fmt.Sprintf("AIScore Filter: RF-балл %d ниже порога %d", score, pl.filterMinScore))
	}
	return nil
}

// Reserve/Unreserve — точка расширения #2 из 4 (была не реализована).
// Оставлены как структурированное логирование lifecycle-события, а не вызов
// внешнего сервиса — намеренно, чтобы не вводить недостающий, непроверенный
// endpoint в ml-server. Ценность уже в самом факте наличия хука: он
// выполняется framework СРАЗУ после выбора ноды, до Bind — естественная
// точка для будущего расширения (например, оповещения ml-server о
// предварительном занятии ресурсов, если понадобится устранить проблему
// устаревших free_cpu/free_mem при высокой конкуренции подов в одном батче).
// Reserve — помимо стандартного lifecycle-логирования, отслеживает
// количество РАЗЛИЧНЫХ нод, на которые уже размещены поды с одним и тем же
// data-path (Направление "уровень репликации" связки Модуль 1 <-> Модуль 2:
// решения AIScore о размещении подов управляют репликацией данных в Alluxio,
// а не отдельный независимый процесс). Если это число достигает
// replicationMinNodes — просит sidecar увеличить минимальное число реплик.
// Ошибка репликации не блокирует Reserve — это оптимизация, не обязательное
// условие успешного размещения.
func (pl *AIScore) Reserve(ctx context.Context, state *framework.CycleState,
	p *v1.Pod, nodeName string) *framework.Status {
	klog.Infof("AIScoreReserve pod=%s node=%s", p.Name, nodeName)

	path, ok := p.Annotations[pl.localityAnnotationKey]
	if !ok || path == "" || pl.replicationURL == "" {
		return nil
	}

	distinctNodes := pl.recordPodPlacement(path, nodeName)
	klog.Infof("AIScoreReplication pod=%s node=%s path=%s distinct_nodes=%d threshold=%d",
		p.Name, nodeName, path, distinctNodes, pl.replicationMinNodes)

	if distinctNodes < pl.replicationMinNodes {
		return nil
	}

	if err := pl.callReplicationService(ctx, path, distinctNodes); err != nil {
		klog.Warningf("AIScoreReplication failed: path=%s replicas=%d err=%v",
			path, distinctNodes, err)
		return nil
	}

	klog.Infof("AIScoreReplication applied: path=%s replicas=%d (triggered by pod=%s node=%s)",
		path, distinctNodes, p.Name, nodeName)
	return nil
}

func (pl *AIScore) Unreserve(ctx context.Context, state *framework.CycleState,
	p *v1.Pod, nodeName string) {
	klog.Infof("AIScoreUnreserve pod=%s node=%s", p.Name, nodeName)
}

// Permit — точка расширения #3 из 4 (была не реализована). Архитектурно
// более чистая реализация fallback-логики Раздела 6.9, чем пост-хок оверрайд
// баллов в NormalizeScore: вместо подмены Score-значений здесь framework
// приостанавливает Bind (Status=Wait) на permitWaitDuration, если calibrated
// confidence для выбранной ноды ниже порога — даёт время на пересмотр
// решения (например, повторный Permit после истечения wait), не изменяя
// сами баллы задним числом.
func (pl *AIScore) Permit(ctx context.Context, state *framework.CycleState,
	p *v1.Pod, nodeName string) (*framework.Status, time.Duration) {

	data, err := state.Read(confidenceStateKey)
	if err != nil {
		// NormalizeScore не записал confidence (например, <2 кандидатов
		// или ошибка calibrate-сервиса) — не блокируем Bind
		return framework.NewStatus(framework.Success), 0
	}
	csd, ok := data.(*confidenceStateData)
	if !ok {
		return framework.NewStatus(framework.Success), 0
	}

	if csd.confidence >= pl.confidenceThreshold {
		pl.clearPermitAttempts(podKey(p)) // успех — сбрасываем счётчик на будущее
		return framework.NewStatus(framework.Success), 0
	}

	attempts := pl.incrementPermitAttempts(podKey(p))
	if attempts > pl.maxPermitRetries {
		klog.Warningf("AIScorePermit pod=%s node=%s confidence=%.4f threshold=%.4f "+
			"превышен лимит попыток (%d) — Unschedulable вместо бесконечного Wait",
			p.Name, nodeName, csd.confidence, pl.confidenceThreshold, pl.maxPermitRetries)
		pl.clearPermitAttempts(podKey(p))
		return framework.NewStatus(framework.Unschedulable,
			fmt.Sprintf("AIScore: calibrated confidence %.4f ниже порога %.4f после %d попыток",
				csd.confidence, pl.confidenceThreshold, pl.maxPermitRetries)), 0
	}

	klog.Infof("AIScorePermit pod=%s node=%s confidence=%.4f threshold=%.4f attempt=%d/%d -> Wait %s",
		p.Name, nodeName, csd.confidence, pl.confidenceThreshold, attempts, pl.maxPermitRetries, pl.permitWaitDuration)
	return framework.NewStatus(framework.Wait), pl.permitWaitDuration
}

func (pl *AIScore) incrementPermitAttempts(key string) int {
	pl.permitMu.Lock()
	defer pl.permitMu.Unlock()
	pl.permitAttempts[key]++
	return pl.permitAttempts[key]
}

func (pl *AIScore) clearPermitAttempts(key string) {
	pl.permitMu.Lock()
	defer pl.permitMu.Unlock()
	delete(pl.permitAttempts, key)
}

func (pl *AIScore) Score(ctx context.Context, state *framework.CycleState,
	pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	nodeInfo, err := pl.handle.SnapshotSharedLister().NodeInfos().Get(nodeName)
	if err != nil {
		return 0, framework.AsStatus(err)
	}

	features := extractFeatures(pod, nodeInfo)
	start := time.Now()
	score, err := pl.callMLService(ctx, features)
	elapsed := time.Since(start)
	klog.Infof("AIScore inference latency: node=%s pod=%s elapsed_ms=%.2f transport=http",
		nodeName, pod.Name, elapsed.Seconds()*1000)

	if err != nil {
		klog.Warningf("AIScore fallback: node=%s pod=%s err=%v elapsed_ms=%.2f transport=http",
			nodeName, pod.Name, err, elapsed.Seconds()*1000)
		return 50, nil
	}
	// Locality bonus (Направление A связки с Модулем 2, Alluxio): если под
	// аннотирован путём к данным и эти данные уже закэшированы на этой
	// ноде — прибавляем бонус к ML-баллу до финального клэмпинга, чтобы
	// бонус участвовал и в итоговом score, и в расчёте margin в
	// NormalizeScore() на общих основаниях с обычным ML-баллом.
	if path, ok := pod.Annotations[pl.localityAnnotationKey]; ok && path != "" {
		lr, lerr := pl.callLocalityService(ctx, path)
		if lerr != nil {
			klog.Warningf("AIScoreLocality lookup failed: node=%s pod=%s path=%s err=%v",
				nodeName, pod.Name, path, lerr)
		} else if nodeMatchesLocality(nodeInfo, lr) {
			klog.Infof("AIScoreLocality bonus applied: node=%s pod=%s path=%s tiers=%v bonus=%d",
				nodeName, pod.Name, path, lr.Tiers, pl.localityBonus)
			score += pl.localityBonus
		}
	}
	if score > framework.MaxNodeScore {
		score = framework.MaxNodeScore
	}
	if score < framework.MinNodeScore {
		score = framework.MinNodeScore
	}
	return score, nil
}



// NormalizeScore вызывается framework один раз НА ПОД, уже после того как
// Score() отработал для ВСЕХ candidate-нод — единственное место в контракте
// ScorePlugin, где можно честно посчитать margin между кандидатами (margin
// физически недоступен внутри одного вызова Score(), который видит только
// одну ноду). Раздел 6.9 мастер-отчёта: margin считается по RF-баллам
// (уже посчитанным и клэмпленным в Score()), конвертируется в calibrated
// confidence через /calibrate_confidence, и при низкой уверенности баллы
// ВСЕХ кандидатов этого пода заменяются на эвристику (fallback).
//
// Дополнительно (для Permit): посчитанная confidence записывается в
// CycleState под ключом confidenceStateKey — Permit() читает её позже в
// конвейере framework и, при необходимости, приостанавливает Bind.
func (pl *AIScore) NormalizeScore(ctx context.Context, state *framework.CycleState,
	pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {

	if len(scores) < 2 {
		return nil // margin не определён для <2 кандидатов — доверяем RF как есть
	}

	sorted := make([]int64, len(scores))
	for i, s := range scores {
		sorted[i] = s.Score
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] > sorted[j] })
	margin := float64(sorted[0] - sorted[1])

	cr, err := pl.callCalibrateService(ctx, margin)
	if err != nil {
		klog.Warningf("AIScore calibrate fallback: pod=%s err=%v — доверяем RF-баллу без учёта confidence",
			pod.Name, err)
		return nil
	}

	state.Write(confidenceStateKey, &confidenceStateData{
		confidence: cr.Confidence,
		calibrated: cr.Calibrated,
	})

	fallbackTriggered := cr.Confidence < pl.confidenceThreshold

	// Структурированное логирование в формате, аналогичном JSONL-логам
	// run_aiscore_over_scenarios.py / compare_to_oracle.py — позволяет
	// парсить логи живого кластера той же аналитикой, что и offline-пилот
	// (Раздел 6.9), не переписывая её под живой формат заново.
	klog.Infof("AIScoreConfidence pod=%s margin=%.4f confidence=%.4f calibrated=%v threshold=%.4f fallback=%v",
		pod.Name, margin, cr.Confidence, cr.Calibrated, pl.confidenceThreshold, fallbackTriggered)

	if !fallbackTriggered {
		return nil
	}

	for i, s := range scores {
		nodeInfo, err := pl.handle.SnapshotSharedLister().NodeInfos().Get(s.Name)
		if err != nil {
			klog.Warningf("AIScore fallback: не удалось получить nodeInfo для %s: %v", s.Name, err)
			continue // оставляем исходный RF-балл для этой ноды, если не смогли получить nodeInfo
		}
		scores[i].Score = heuristicScore(nodeInfo)
	}
	return nil
}

func (pl *AIScore) ScoreExtensions() framework.ScoreExtensions {
	return pl
}

func New(ctx context.Context, _ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	modelURL := os.Getenv("ML_SERVER_URL")
	if modelURL == "" {
		modelURL = "http://ml-server:8000/predict"
	}

	calibrateURL := os.Getenv("ML_CALIBRATE_URL")
	if calibrateURL == "" {
		calibrateURL = "http://ml-server:8000/calibrate_confidence"
	}

	// ВАЖНО (Раздел 6.9.4 мастер-отчёта): после калибровки confidence может
	// сжаться в узкий диапазон (например, [0.65, 1.0]) — абсолютный порог,
	// подобранный "на глаз" до калибровки, ломается. Значение по умолчанию
	// ниже — временное и НЕ основано на процентиле реального распределения
	// confidence; перед боевым использованием пересчитайте его как, например,
	// 20-й процентиль confidence на калибровочном сплите (Раздел 6.9.4) и
	// зафиксируйте через ML_CONFIDENCE_THRESHOLD, а не полагайтесь на дефолт.
	confidenceThreshold := 0.7
	if v := os.Getenv("ML_CONFIDENCE_THRESHOLD"); v != "" {
		if parsed, err := strconv.ParseFloat(v, 64); err == nil {
			confidenceThreshold = parsed
		} else {
			klog.Warningf("AIScore: не удалось распарсить ML_CONFIDENCE_THRESHOLD=%q, используется %.2f",
				v, confidenceThreshold)
		}
	}

	// filterMinScore: -1 = Filter выключен (дефолт, консервативно).
	filterMinScore := int64(-1)
	if v := os.Getenv("ML_FILTER_MIN_SCORE"); v != "" {
		if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
			filterMinScore = parsed
		} else {
			klog.Warningf("AIScore: не удалось распарсить ML_FILTER_MIN_SCORE=%q, Filter остаётся выключен", v)
		}
	}
	localityURL := os.Getenv("ML_LOCALITY_URL")
	// Пусто по умолчанию — locality bonus выключен, пока явно не задан URL
	// sidecar'а; намеренно консервативный дефолт, как и filterMinScore выше.

	localityBonus := int64(15)
	if v := os.Getenv("ML_LOCALITY_BONUS"); v != "" {
		if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
			localityBonus = parsed
		} else {
			klog.Warningf("AIScore: не удалось распарсить ML_LOCALITY_BONUS=%q, используется %d",
				v, localityBonus)
		}
	}

	localityAnnotationKey := os.Getenv("ML_LOCALITY_ANNOTATION_KEY")
	if localityAnnotationKey == "" {
		localityAnnotationKey = "alluxio.io/data-path"
	}

	replicationURL := os.Getenv("ML_REPLICATION_URL")
	// Пусто по умолчанию — репликация выключена, пока явно не задан URL.

	replicationMinNodes := 2
	if v := os.Getenv("ML_REPLICATION_MIN_NODES"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			replicationMinNodes = parsed
		} else {
			klog.Warningf("AIScore: не удалось распарсить ML_REPLICATION_MIN_NODES=%q, используется %d",
				v, replicationMinNodes)
		}
	}
	permitWaitDuration := 3 * time.Second
	if v := os.Getenv("ML_PERMIT_WAIT_SECONDS"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			permitWaitDuration = time.Duration(parsed) * time.Second
		} else {
			klog.Warningf("AIScore: не удалось распарсить ML_PERMIT_WAIT_SECONDS=%q, используется %s",
				v, permitWaitDuration)
		}
	}

	maxPermitRetries := 3
	if v := os.Getenv("ML_MAX_PERMIT_RETRIES"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			maxPermitRetries = parsed
		} else {
			klog.Warningf("AIScore: не удалось распарсить ML_MAX_PERMIT_RETRIES=%q, используется %d",
				v, maxPermitRetries)
		}
	}

	return &AIScore{
		handle:              h,
		httpClient:          &http.Client{Timeout: 2 * time.Second},
		modelURL:            modelURL,
		calibrateURL:        calibrateURL,
		confidenceThreshold: confidenceThreshold,
		filterMinScore:      filterMinScore,
		permitWaitDuration:  permitWaitDuration,
		maxPermitRetries:    maxPermitRetries,
		permitAttempts:      make(map[string]int),
		localityURL:            localityURL,
		localityBonus:          localityBonus,
		localityAnnotationKey:  localityAnnotationKey,
		localityHTTPClient:     &http.Client{Timeout: 10 * time.Second},
		replicationURL:         replicationURL,
		replicationMinNodes:    replicationMinNodes,
		pathNodes:              make(map[string]map[string]bool),
		replicationHTTPClient:  &http.Client{Timeout: 10 * time.Second},
	}, nil
}
