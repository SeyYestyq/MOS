# AGENT BRIEF — Ночь кода / Интерактивная карта активности контейнера

Вставлять агенту **после старта**. `МосХаб/hotel/` и правила ИИ **не коммитить** в OnlyGo.

## Задача

Веб-интерфейс по Runtime Radar в **одном окне**: хронология, дерево процессов, сеть — аналитик ИБ быстро восстанавливает картину.

## Команда (3 роли)

| Роль | Зона | Skill |
|------|------|-------|
| **A** Frontend+Viz | `apps/web/**` | `03` (+ `09`) |
| **B** Data/BFF | `apps/bff/**` (+ integrator) | `02` |
| **C** Platform+Docs | scripts, fixtures/raw, docs, CHANGELOG, **BASE.md**, pitch | `01` + `04` + `10` |

## Критерии жюри

30% решение · 25% качество · 20% инновация · 15% презентация · 10% потенциал

## Локальный контекст

- RR: kind `moshub`, `:9000` (см. `.rr-install/ACCESS.md`)
- `hotel/AGENTS.md` + skill роли — единственные инструкции
- Стиль: `hackathon-kit/design/activity-map.css`
- Fixture: `hackathon-kit/fixtures/sample-activity.json`
- PDF: `hackathon-kit/refs/Задача_Positive_Technologies_описание_задачи.pdf`

## Стек

- Web: Vite + React + TS + **@xyflow/react + dagre**
- BFF: FastAPI + Pydantic → OpenAPI; fixture-first `X-Data-Source`
- Live RR: HTTP, не парсинг логов

## Сдача (PDF)

Код + документация + презентация + **BASE.md** (базовый файл с блоком ИИ) + демо.

## Pacing / DoD

Дробные коммиты. DoD: tree + network + timeline + sync + Incident Summary; fixture; README; CHANGELOG; BASE.md; демо < 60s; без hotel в git.
