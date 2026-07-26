#!/usr/bin/env bash
## 2. Однородная/гетерогенная нагрузка (Раздел 7ter)
mkdir -p results_heuristic
create_nodes() {
    declare -A specs=(["8"]="16Gi" ["16"]="32Gi" ["24"]="48Gi")
    for cpu in 8 16 24; do
        mem="${specs[$cpu]}"

        response=$(curl -s -X POST http://localhost:3131/api/v1/nodes \
          -H "Content-Type: application/json" \
          -d '{"apiVersion":"v1","kind":"Node","metadata":{"generateName":"node-"}}')

        node_name=$(echo "$response" | python3 -c "import json,sys; print(json.load(sys.stdin)['metadata']['name'])" 2>/dev/null)

        if [ -z "$node_name" ]; then
            echo "ОШИБКА: не удалось создать ноду (cpu=$cpu). Ответ API: $response"
            continue
        fi

        curl -s -X PUT "http://localhost:3131/api/v1/nodes/${node_name}/status" \
          -H "Content-Type: application/json" \
          -d "{\"apiVersion\":\"v1\",\"kind\":\"Node\",\"metadata\":{\"name\":\"${node_name}\"},\"status\":{\"allocatable\":{\"cpu\":\"$cpu\",\"memory\":\"$mem\",\"pods\":\"110\"},\"capacity\":{\"cpu\":\"$cpu\",\"memory\":\"$mem\",\"pods\":\"110\"},\"conditions\":[{\"type\":\"Ready\",\"status\":\"True\"}]}}" > /dev/null

        curl -s -X PATCH "http://localhost:3131/api/v1/nodes/${node_name}" \
          -H "Content-Type: application/strategic-merge-patch+json" \
          -d '{"spec":{"taints":[]}}' > /dev/null

        echo "Создана нода $node_name: cpu=$cpu, memory=$mem"
    done
    sleep 5
}

#create_nodes
#python3 benchmark/run_benchmark.py --label heuristic_heterogeneous --n 8 --out-dir results_heuristic

docker compose -f compose.yml -f compose.local.yml down -v
docker compose -f compose.yml -f compose.local.yml up -d

until curl -s -f http://localhost:3131/api/v1/nodes > /dev/null 2>&1; do
    echo "Жду готовности API..."
    sleep 2
done

create_nodes

# Проверка — сколько реально создано
curl -s http://localhost:3131/api/v1/nodes | python3 -c "
import json, sys
d = json.load(sys.stdin)
print(f'Реально создано нод: {len(d[\"items\"])}')
"

#python3 benchmark/run_benchmark.py --label heuristic_homogeneous --n 8 --homogeneous medium --out-dir results_heuristic
