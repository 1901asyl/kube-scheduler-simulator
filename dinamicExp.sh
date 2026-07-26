mkdir -p results_dynamic

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

# 1. Извлечь паттерн прибытия из реального трейса
#python3 data/extract_arrival_pattern.py

# 2. Чистый кластер с достаточным числом нод (динамическая нагрузка может создать
#    больше "волн" на одних и тех же ресурсах, чем статический 40-подовый прогон)
docker compose -f compose.yml -f compose.local.yml down -v
docker compose -f compose.yml -f compose.local.yml up -d
sleep 8
create_nodes

# 3. Прогнать динамическую нагрузку
python3 benchmark/run_dynamic_benchmark.py --label dynamic_trace --out-dir results_dynamic

