# Scripts — ветка Дамира

| Скрипт | Зачем |
|--------|--------|
| `gen-activity.sh` | 4 фазы + nc:8888 + flush `ls` |
| `dump-rr-events.sh` | REST slice → `fixtures/raw/events.json` (с `args`) |

```bash
# после Sources/Allow/All
export KUBE_CONTEXT=kind-moshub   # если нужно
./scripts/gen-activity.sh

export RR_BASE_URL=https://localhost:9000
export RR_TOKEN=...               # Access-Token из UI
./scripts/dump-rr-events.sh

cd apps/bff
DATA_MODE=raw RAW_DUMP_PATH=../../fixtures/raw/events.json go run .
```

Go BFF: `../apps/bff` — **Go**, fixture-first; live через тот же slice API.
