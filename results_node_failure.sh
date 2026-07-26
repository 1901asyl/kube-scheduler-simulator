#!/usr/bin/env bash
mkdir -p results_node_failure

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





# Чистый кластер с ЗАПАСОМ ёмкости (важно! иначе отказ узла сразу приведёт
# к массовым провалам просто из-за общей нехватки ресурсов, как в 7quater
# с 3 нодами — тут же нас интересует ИМЕННО эффект отказа, а не нехватку ёмкости изначально)
docker compose -f compose.yml -f compose.local.yml down -v
docker compose -f compose.yml -f compose.local.yml up -d
sleep 8
create_nodes_large   # 4 ноды, 80 CPU суммарно — та же ёмкость, что дала 100/100 в 7quater

python3 benchmark/run_dynamic_benchmark_with_failure.py \
  --label node_failure_early_busiest \
  --out-dir results_node_failure \
  --failure-fraction 0.3 2>&1 | tee results_node_failure/node_failure_early_busiest.log
