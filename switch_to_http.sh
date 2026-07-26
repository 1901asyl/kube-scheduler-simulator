#!/usr/bin/env bash
# Переключает рабочую копию на HTTP-версию (versions/http/ -> рабочие файлы).
set -euo pipefail
cd "$(dirname "$0")"

cp versions/http/plugin.go simulator/docs/sample/ai_score/plugin.go
cp versions/http/main.py ml-server/main.py
cp versions/http/requirements.txt ml-server/requirements.txt
cp versions/http/Dockerfile ml-server/Dockerfile

echo "Переключено на HTTP. Порт ml-server должен быть 8000 — проверьте compose.local.yml вручную:"
grep -n "8000\|5555" compose.local.yml || true
echo "Проверка: grep transport= в рабочем plugin.go должен показать transport=http"
grep -n "transport=" simulator/docs/sample/ai_score/plugin.go
