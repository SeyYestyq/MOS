# AGENTS.md (локально / hotel — НЕ коммитить в OnlyGo)

Хакатон: Ночь кода · задача PT · Activity Map на Runtime Radar.  
Команда: **3 человека** (A / B / C).

## Стек (фиксировано)

- `apps/web`: Vite + React + TypeScript + **@xyflow/react** + layout (dagre / `@dagrejs/dagre`)
- `apps/bff`: FastAPI + Pydantic модели = источник правды контракта (`/openapi.json`)
- Данные: Runtime Radar HTTP API **или** fixture JSON
- Стиль: `../hackathon-kit/design/activity-map.css` (один импорт; не искать CSS вне kit)

## Path ownership

| Роль | Можно писать | Нельзя |
|------|--------------|--------|
| **A** Frontend+Viz | `apps/web/**` | `apps/bff/**`, корень, docs сдачи |
| **B** Data/BFF | `apps/bff/**`; корень как integrator | `apps/web/**` без joint |
| **C** Platform+Docs | `scripts/**`, `deploy/**`, `fixtures/raw/**`, `docs/**`, `CHANGELOG.md`, `BASE.md`, `README.md`, pptx | `apps/**` без просьбы |

**Корень репозитория** (выше `apps/`, кроме явно отданных C файлов docs/CHANGELOG/BASE/README) — **B** (integrator).  
C пишет `CHANGELOG.md`, `BASE.md`, `README.md`, `docs/**` без переписывания package/docker без B.

В каждом промпте агенту первой строкой:  
`Разрешённые пути: … . Всё остальное read-only. Не меняй .gitignore, docker-compose, Makefile без явной просьбы.`

## Контракт

Источник правды: Pydantic в BFF → `/openapi.json`.  
Breaking change — только на sync команды (ведёт B).

## Handoff

Предпочтительно: OpenAPI / curl + AC в чате.  
`hotel/handoffs/` — опционально.

## Pacing коммитов в OnlyGo

Не mega-commit через 15 минут. Дробить:  
scaffold → fixture → tree → timeline/network → sync → summary → live → docs/BASE.

## Skills

Только `hotel/skills/` (карта — `hotel/SKILLS.md`).

| Роль | Файл |
|------|------|
| A | `skills/03-frontend-activity-map.md` (+ `09` для Incident Summary) |
| B | `skills/02-backend-activity-bff.md` |
| C | `01-platform-rr.md`, `04-pitch-demo.md`, `10-docs-changelog.md`, `09-soc-demo-language.md` |
| Все | `08-git-secrets.md` |
| Freeze | `05-freeze-review.md` |
| Стык | `06-handoff-acid.md` |
| Баг | `07-focused-fix.md` |
| Опц. | `skills/_borghei/` — не грузить пачкой |

## DoD продукта

Одно окно: process tree + network + timeline + sync; Incident Summary; fixture fallback; README; **CHANGELOG**; **BASE.md** (ИИ); демо < 60s; Public.
