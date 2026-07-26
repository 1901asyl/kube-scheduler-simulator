#!/usr/bin/env bash
set -e
mkdir -p results_overhead

SCHEDULER_YAML="simulator/cmd/scheduler/scheduler.yaml"

create_nodes_large() {
    declare -A specs=(["8"]="16Gi" ["16"]="32Gi" ["24"]="48Gi" ["32"]="64Gi")
    for cpu in 8 16 24 32; do
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

        echo "Создана нода: cpu=$cpu, memory=$mem"
    done
    sleep 5
}

cat > scheduler_baseline.yaml << 'YAML'
apiVersion: kubescheduler.config.k8s.io/v1
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: "/etc/kube-scheduler-simulator/kubeconfig.yaml"
profiles:
  - schedulerName: default-scheduler
YAML

cat > scheduler_aiscore.yaml << 'YAML'
apiVersion: kubescheduler.config.k8s.io/v1
kind: KubeSchedulerConfiguration
clientConnection:
  kubeconfig: "/etc/kube-scheduler-simulator/kubeconfig.yaml"
profiles:
  - schedulerName: default-scheduler
    plugins:
      score:
        enabled:
          - name: AIScore
            weight: 10
YAML

run_config() {
    local label=$1
    local scheduler_yaml=$2
    local noop_mode=$3

    echo "=== Конфигурация: ${label} (scheduler=${scheduler_yaml}, NOOP_MODE=${noop_mode}) ==="

    cp "$scheduler_yaml" "$SCHEDULER_YAML"
    sed -i "s/NOOP_MODE=.*/NOOP_MODE=${noop_mode}/" compose.local.yml
    sed -i "s/ARTIFICIAL_DELAY_MS=.*/ARTIFICIAL_DELAY_MS=0/" compose.local.yml

    docker compose -f compose.yml -f compose.local.yml down -v
    docker compose -f compose.yml -f compose.local.yml build simulator-scheduler
    docker compose -f compose.yml -f compose.local.yml up -d
    sleep 8

    create_nodes_large

    python3 benchmark/run_benchmark.py --label "$label" --n 8 --out-dir results_overhead \
      2>&1 | tee "results_overhead/${label}.log"
}

run_config "overhead_baseline"     "scheduler_baseline.yaml" "0"
run_config "overhead_aiscore_real" "scheduler_aiscore.yaml"  "0"
run_config "overhead_aiscore_noop" "scheduler_aiscore.yaml"  "1"

echo ""
echo "=== СВОДКА: per-pod scheduling latency ==="
for label in overhead_baseline overhead_aiscore_real overhead_aiscore_noop; do
    echo "--- ${label} ---"
    grep "Per-pod scheduling latency" "results_overhead/${label}.log" || echo "(нет данных о latency)"
done
