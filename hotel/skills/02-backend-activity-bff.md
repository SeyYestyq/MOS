# Backend Activity BFF

Всё нужное здесь. Не открывать senior-backend целиком.

## Зона

`apps/bff/**` + integrator корня. Web — read-only.

## Стек

- **Go (chi/net/http) или FastAPI** — выбрать на T0; RR-продукт на Go, но BFF может быть любым
- Контракт JSON: процессы с **`exec_id`**, **`parent_exec_id`**, pid/binary/args; сеть; events
- Нормализатор dump Tetragon/RR → nodes/edges для React Flow + items для vis-timeline
- `GET /api/activity?pod=&namespace=` → payload
- `X-Data-Source: fixture | live`
- Live: public-api через reverse-proxy (`:9000` / `:32000`), не парсинг логов
- Тест: fixture валидируется / golden JSON

## Порядок коммитов (pacing)

1. Каркас app + роут health/activity stub  
2. Pydantic-модели + чтение fixture  
3. Нормализация dump → ActivityPayload  
4. Live client (опционально) + fallback на fixture  
5. pytest схемы  

## Паттерны API

- Явные типы, без «сырого dict» в публичном ответе.
- Ошибки RR → 200 + fixture + `X-Data-Source: fixture` (демо не должно краснеть) или 502 только если явно `?source=live`.
- Без БД, без auth-сервиса, без GraphQL.
- CORS только для Vite origin.

## Контракт (мини-чеклист API)

- Один ресурс activity, поля стабильны (Pydantic).
- Ошибки RR не валят демо (fallback fixture).
- OpenAPI отдаётся; breaking rename — только на sync.
- Точечно: `_borghei/api-design-reviewer.md`, pytest идеи `_borghei/api-test-suite-builder.md`

## Не делать

Миграции БД, load-test, Express scaffolder, микросервисный сплит.
