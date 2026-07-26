#!/usr/bin/env bash

mkdir -p results_scenarios

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

# 1. Вернуть нормальные условия (снять лимит памяти и задержку с прошлых экспериментов)
sed -i "s/mem_limit:.*/mem_limit: 1g/" compose.local.yml
sed -i "s/ARTIFICIAL_DELAY_MS=.*/ARTIFICIAL_DELAY_MS=0/" compose.local.yml

# 2. Чистый кластер
docker compose -f compose.yml -f compose.local.yml down -v
docker compose -f compose.yml -f compose.local.yml up -d
sleep 8

# 3. Те же 3 гетерогенные ноды (переиспользуем функцию create_nodes из предыдущих скриптов)
create_nodes

# 4. Однородная нагрузка — 40 одинаковых подов профиля medium
python3 benchmark/run_benchmark.py --label homogeneous_medium --n 8 --homogeneous medium --out-dir results_scenarios

# 5. Для сравнения — гетерогенный прогон в тех же условиях (свежий кластер)
docker compose -f compose.yml -f compose.local.yml down -v
docker compose -f compose.yml -f compose.local.yml up -d
sleep 8
create_nodes
python3 benchmark/run_benchmark.py --label heterogeneous_mixed --n 8 --out-dir results_scenarios

# 6. Сравнение
python3 benchmark/run_benchmark.py --compare homogeneous_medium heterogeneous_mixed --out-dir results_scenarios
