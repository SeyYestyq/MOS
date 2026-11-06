# AGENTS.md (локально / hotel — НЕ коммитить в OnlyGo)

Хакатон: Ночь кода · задача PT · Activity Map на Runtime Radar.

## Стек (фиксировано)

- `apps/web`: Vite + React + TypeScript + **@xyflow/react** + layout (dagre / `@dagrejs/dagre`)
- `apps/bff`: FastAPI + Pydantic модели = источник правды контракта (`/openapi.json`)
- Данные: Runtime Radar HTTP API **или** fixture JSON
- Стиль: `../hackathon-kit/design/activity-map.css` (один импорт; не искать CSS вне kit)

## Path ownership

| Роль | Можно писать | Нельзя |
|------|--------------|--------|
| P1 | `scripts/**`, `deploy/**`, сырые dumps в `fixtures/raw/**` | `apps/**` без просьбы |
| P2 | `apps/bff/**` | `apps/web/**`, корень репо |
| P3 | `apps/web/**` | `apps/bff/**`, корень репо |
| P4 | `docs/pitch/**`, `README.md`, pptx | `apps/**` без joint |
| Integrator (P1 или P2) | корень: `docker-compose.yml`, корневой README merge, package workspaces | — |

**Корень репозитория** (файлы выше `apps/`, `scripts/`, `docs/pitch/`) — только integrator.

В каждом промпте агенту первой строкой:  
`Разрешённые пути: … . Всё остальное read-only. Не меняй .gitignore, docker-compose, Makefile без явной просьбы.`

## Контракт

Источник правды: Pydantic в BFF → `/openapi.json`.  
Markdown CONTRACT — шпаргалка. Breaking change — только на sync команды.

## Handoff

Предпочтительно: «вот OpenAPI / curl + AC в чате».  
Markdown в `hotel/handoffs/` — по желанию, не обязателен каждый час.

## Pacing коммитов в OnlyGo

Не вываливать 40 файлов через 15 минут после старта. Дробить:  
scaffold → fixture read → tree → timeline → sync click → live RR → polish.

## Skills

Только `hotel/skills/` (карта — `hotel/SKILLS.md`).  
Не подключать `claude-skills-main` и чужие skill-репозитории.

| Роль | Файл |
|------|------|
| P1 | `skills/01-platform-rr.md` |
| P2 | `skills/02-backend-activity-bff.md` |
| P3 | `skills/03-frontend-activity-map.md` |
| P4 | `skills/04-pitch-demo.md` |
| Freeze | `skills/05-freeze-review.md` |
| Стык | `skills/06-handoff-acid.md` |

## DoD продукта

Одно окно: process tree + network + timeline + sync; Incident Summary; fixture fallback; README; демо < 60s time-to-insight.
