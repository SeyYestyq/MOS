# apps/bff — Go Activity BFF

Стек: **Go 1.22+**, `net/http`, без внешних модулей.

```bash
cd hackathon-kit/apps/bff
go test ./...
go run .
# GET http://localhost:8000/healthz
# GET http://localhost:8000/api/activity?namespace=hackathon&pod=attacker-pod
# Header: X-Data-Source: fixture | raw | live
```

| DATA_MODE | Источник |
|-----------|----------|
| `fixture` (default) | `fixtures/sample-activity.json` |
| `raw` | dump Дамира `fixtures/raw/events.json` → нормализатор |
| `live` | REST slice + `RR_TOKEN` (fallback → fixture) |

Контракт = `stubs/types.ts`. XSS sanitize на arguments. Рёбра дерева только `parentExecId → execId`.

После старта: скопировать каталог в OnlyGo `apps/bff/`.
