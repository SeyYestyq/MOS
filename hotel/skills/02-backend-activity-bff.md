# Backend Activity BFF (P2)

Всё нужное здесь. Не открывать `claude-skills-main` / senior-backend.

## Зона

Только `apps/bff/**`. Корень / web — read-only.

## Стек

- FastAPI + Pydantic v2
- Модели = контракт: `ProcessNode`, `NetEdge`, `ActivityEvent`, `ActivityPayload`
- `GET /api/activity?pod=&namespace=` → payload
- Заголовок `X-Data-Source: fixture | live`
- Fixture JSON читается с диска; live — HTTP к Runtime Radar (`:9000` + token из env), не парсинг логов
- Один pytest: ответ валидируется схемой Pydantic
- OpenAPI: `/openapi.json` — источник правды для фронта

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

## Не делать

Миграции БД, load-test скрипты, Express/Fastify scaffolder, микросервисный сплит.
