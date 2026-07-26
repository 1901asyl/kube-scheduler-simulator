#!/usr/bin/env bash
mkdir -p results_memory

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

# --- Разовая проверка: реально ли применяется лимит + пиковая память во время нагрузки ---
sed -i "s/mem_limit:.*/mem_limit: 120m/" compose.local.yml
docker compose -f compose.yml -f compose.local.yml down -v
docker compose -f compose.yml -f compose.local.yml up -d
sleep 5

echo "=== Проверка применённого лимита ==="
docker inspect ml-server --format='Limit: {{.HostConfig.Memory}} bytes'

create_nodes

( while true; do
    docker stats ml-server --no-stream --format "{{.MemUsage}}" 2>/dev/null
    sleep 1
done ) > results_memory/live_mem_120m.log &
MONITOR_PID=$!

python3 benchmark/run_benchmark.py --label memtest_check2 --n 8 --out-dir results_memory

kill $MONITOR_PID 2>/dev/null

echo "=== Пиковое использование памяти во время прогона (120m) ==="
sort -t/ -k1 -h results_memory/live_mem_120m.log | tail -5

# --- Основной цикл по всем значениям лимита ---
for MEM_LIMIT in "400m" "250m" "200m" "170m" "150m" "120m"; do
    echo "=== MEM_LIMIT=$MEM_LIMIT ==="
    sed -i "s/mem_limit:.*/mem_limit: $MEM_LIMIT/" compose.local.yml

    docker compose -f compose.yml -f compose.local.yml down -v
    docker compose -f compose.yml -f compose.local.yml up -d
    sleep 8

    create_nodes

    { time python3 benchmark/run_benchmark.py --label "memtest_${MEM_LIMIT}" --n 8 --out-dir results_memory ; } \
      2>&1 | tee "results_memory/memtest_${MEM_LIMIT}.log"

    oom_killed=$(docker inspect ml-server --format='{{.State.OOMKilled}}')
    restart_count=$(docker inspect ml-server --format='{{.RestartCount}}')
    echo "OOMKilled=$oom_killed RestartCount=$restart_count" | tee -a "results_memory/memtest_${MEM_LIMIT}.log"

    docker compose -f compose.yml -f compose.local.yml logs simulator-scheduler | \
      grep "AIScore inference latency" | \
      grep -oP 'elapsed_ms=\K[0-9.]+' > "results_memory/real_latency_${MEM_LIMIT}.txt"
done

echo ""
echo "=== СВОДКА ==="
for MEM_LIMIT in "400m" "250m" "200m" "170m" "150m" "120m"; do
    label="memtest_${MEM_LIMIT}"
    real_time=$(grep "^real" "results_memory/${label}.log" | awk '{print $2}')
    oom=$(grep "OOMKilled=" "results_memory/${label}.log")
    scheduled=$(python3 -c "
import csv
with open('results_memory/${label}_results.csv') as f:
    rows = list(csv.DictReader(f))
    print(f'{sum(1 for r in rows if r[\"scheduled\"]==\"True\")}/{len(rows)}')
" 2>/dev/null || echo "N/A")
    latency_stats=""
    if [ -s "results_memory/real_latency_${MEM_LIMIT}.txt" ]; then
        latency_stats=$(python3 -c "
import numpy as np
data = np.loadtxt('results_memory/real_latency_${MEM_LIMIT}.txt')
print(f'mean={data.mean():.1f}ms p95={np.percentile(data,95):.1f}ms')
" 2>/dev/null)
    fi
    echo "MEM_LIMIT=${MEM_LIMIT}: time=$real_time scheduled=$scheduled $oom latency=[$latency_stats]"
done
