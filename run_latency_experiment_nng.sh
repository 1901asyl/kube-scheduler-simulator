#!/usr/bin/env bash
# Полный свип: транспорт=NNG, NOOP_MODE={0,1} x ARTIFICIAL_DELAY_MS={0,50,100,500,1500}.
# NOOP_MODE=0 -> сеть + сериализация + инференс
# NOOP_MODE=1 -> сеть + сериализация (model.predict() не вызывается)
set -euo pipefail

# --- GUARD: не даём случайно прогнать не тот транспорт под этим именем ---
PLUGIN_FILE="simulator/docs/sample/ai_score/plugin.go"
if grep -q "transport=http" "$PLUGIN_FILE" 2>/dev/null; then
    echo "ОШИБКА: в $PLUGIN_FILE сейчас HTTP-версия (найдено 'transport=http')," >&2
    echo "а этот скрипт помечен как NNG-эксперимент. Выполните ./switch_to_nng.sh перед запуском." >&2
    exit 1
fi
if ! grep -q "transport=nng" "$PLUGIN_FILE" 2>/dev/null; then
    echo "ПРЕДУПРЕЖДЕНИЕ: не найдена метка 'transport=nng' в $PLUGIN_FILE — " >&2
    echo "убедитесь, что рабочая копия действительно NNG-версия." >&2
fi

# --- Пересборка ДО начала свипа, с остановкой при ошибке (set -e это обеспечит) ---
# Раньше скрипт делал только `up -d`, что при неудачной сборке молча
# запускало СТАРЫЙ образ — итог был неотличим от "всё ок", кроме пустых
# real_latency_*.txt. Теперь сборка — явный шаг, и её падение сразу
# останавливает весь эксперимент.
echo "=== Пересборка образов (может занять время) ==="
docker compose -f compose.yml -f compose.local.yml build ml-server simulator-scheduler

mkdir -p results_timing_nng

create_nodes() {
    declare -A specs=(["8"]="16Gi" ["16"]="32Gi" ["24"]="48Gi")
    for cpu in 8 16 24; do
        mem="${specs[$cpu]}"

        node_name=$(curl -s -X POST http://localhost:3131/api/v1/nodes \
          -H "Content-Type: application/json" \
          -d '{"apiVersion":"v1","kind":"Node","metadata":{"generateName":"node-"}}' \
          | python3 -c "import json,sys; print(json.load(sys.stdin)['metadata']['name'])")

        curl -s -X PUT "http://localhost:3131/api/v1/nodes/${node_name}/status" \
          -H "Content-Type: application/json" \
          -d "{\"apiVersion\":\"v1\",\"kind\":\"Node\",\"metadata\":{\"name\":\"${node_name}\"},\"status\":{\"allocatable\":{\"cpu\":\"$cpu\",\"memory\":\"$mem\",\"pods\":\"110\"},\"capacity\":{\"cpu\":\"$cpu\",\"memory\":\"$mem\",\"pods\":\"110\"},\"conditions\":[{\"type\":\"Ready\",\"status\":\"True\"}]}}" > /dev/null

        curl -s -X PATCH "http://localhost:3131/api/v1/nodes/${node_name}" \
          -H "Content-Type: application/strategic-merge-patch+json" \
          -d '{"spec":{"taints":[]}}' > /dev/null
    done
    sleep 5
}

for NOOP in 0 1; do
  for DELAY in 0 50 100 500 1500; do
    label="timing_nng_noop${NOOP}_${DELAY}ms"
    echo "=== [NNG] NOOP_MODE=$NOOP ARTIFICIAL_DELAY_MS=$DELAY ==="

    sed -i "s/ARTIFICIAL_DELAY_MS=.*/ARTIFICIAL_DELAY_MS=$DELAY/" compose.local.yml
    sed -i "s/NOOP_MODE=.*/NOOP_MODE=$NOOP/" compose.local.yml

    docker compose -f compose.yml -f compose.local.yml down -v
    docker compose -f compose.yml -f compose.local.yml up -d
    sleep 8

    # --- Runtime-проверка: плагин реально инициализировался? ---
    # Ловит случай, когда build прошёл, но New() упал (например, не смог
    # законнектиться к ml-server) — тогда обычный scheduler мог продолжить
    # работу с другим плагином/дефолтным скорингом молча.
    if ! docker compose -f compose.yml -f compose.local.yml logs simulator-scheduler | grep -q "AIScore"; then
        echo "ОШИБКА: в логах simulator-scheduler нет ни одной строки 'AIScore' —" >&2
        echo "плагин не инициализировался. Проверьте: docker compose logs simulator-scheduler" >&2
        exit 1
    fi

    create_nodes

    { time python3 benchmark/run_benchmark.py --label "$label" --n 8 --out-dir results_timing_nng ; } \
      2>&1 | tee "results_timing_nng/${label}.log"

    docker compose -f compose.yml -f compose.local.yml logs simulator-scheduler | \
      grep "AIScore inference latency" | \
      grep -oP 'elapsed_ms=\K[0-9.]+' > "results_timing_nng/real_latency_${label}.txt"
  done
done

echo ""
echo "=== СВОДКА (NNG) ==="
for NOOP in 0 1; do
  for DELAY in 0 50 100 500 1500; do
    label="timing_nng_noop${NOOP}_${DELAY}ms"
    real_time=$(grep "^real" "results_timing_nng/${label}.log" | awk '{print $2}')
    latency_stats=""
    if [ -s "results_timing_nng/real_latency_${label}.txt" ]; then
        latency_stats=$(python3 -c "
import numpy as np
data = np.loadtxt('results_timing_nng/real_latency_${label}.txt')
print(f'mean={data.mean():.2f}ms p95={np.percentile(data,95):.2f}ms')
")
    fi
    echo "NOOP=${NOOP} DELAY=${DELAY}ms: time=$real_time real_latency=[$latency_stats]"
  done
done
