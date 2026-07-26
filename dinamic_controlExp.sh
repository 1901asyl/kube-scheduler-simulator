#!/usr/bin/env bash
mkdir -p results_dynamic_control

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



# Пересобрать чистый кластер
docker compose -f compose.yml -f compose.local.yml down -v
docker compose -f compose.yml -f compose.local.yml up -d
sleep 8

create_nodes_large

echo "=== Проверка ёмкости кластера ==="
curl -s http://localhost:3131/api/v1/nodes | python3 -c "
import json, sys
d = json.load(sys.stdin)
total_cpu = sum(int(n['status']['allocatable']['cpu']) for n in d['items'])
print(f'Нод: {len(d[\"items\"])}, суммарный CPU: {total_cpu}')
"

echo "=== Запуск динамической нагрузки (та же, что и в контрольной группе) ==="
python3 benchmark/run_dynamic_benchmark.py --label dynamic_trace_large --out-dir results_dynamic_control
