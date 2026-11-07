# Backend Activity BFF (Go)

Всё нужное здесь.

## Зона

`apps/bff/**` + integrator корня. Web — read-only.

## Стек (зафиксировано)

- **Go 1.22+**, stdlib `net/http` (опционально `go-chi/chi` если удобно)
- Контракт JSON: процессы с **`exec_id`**, **`parent_exec_id`**, pid/binary/args; сеть; events
- Нормализатор dump Tetragon/RR → nodes/edges для React Flow + items для vis-timeline
- `GET /healthz`, `GET /api/activity?pod=&namespace=`
- Заголовок ответа `X-Data-Source: fixture | live`
- Live: public-api через reverse-proxy (`:9000` / `:32000`), не парсинг логов
- Тест: `go test` на normalize + чтение fixture
- Старт с stub: `hackathon-kit/stubs/bff/main.go` + `go.mod`
- Env: `hackathon-kit/stubs/bff/env.example` → локально `.env` (не в OnlyGo)
- Dump от **Platform** в `fixtures/raw/`

## Порядок коммитов (pacing)

1. `go mod` + `/healthz` + stub `/api/activity`  
2. structs + чтение fixture  
3. Normalize dump → ActivityPayload  
4. Live HTTP client + fallback fixture  
5. `go test` на normalize  

## Паттерны API

- Явные struct + `encoding/json`, без «сырого map» в публичном ответе без нужды.
- Ошибки RR → 200 + fixture + `X-Data-Source: fixture` (демо не краснеет).
- Без БД, без auth-сервиса, без GraphQL.
- CORS только Vite origin.

## Не делать

FastAPI/Python BFF (стек = Go); миграции БД; микросервисный сплит.
