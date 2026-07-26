#!/usr/bin/env bash
# Переключает рабочую копию на NNG-версию (versions/nng/ -> рабочие файлы).
set -euo pipefail
cd "$(dirname "$0")"

cp versions/nng/plugin.go simulator/docs/sample/ai_score/plugin.go
cp versions/nng/main.py ml-server/main.py
cp versions/nng/requirements.txt ml-server/requirements.txt
cp versions/nng/Dockerfile ml-server/Dockerfile

echo "Переключено на NNG. Порт ml-server должен быть 5555 — проверьте compose.local.yml вручную:"
grep -n "8000\|5555" compose.local.yml || true
echo "Проверка: grep transport= в рабочем plugin.go должен показать transport=nng"
grep -n "transport=" simulator/docs/sample/ai_score/plugin.go
