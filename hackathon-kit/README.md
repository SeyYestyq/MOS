# Hackathon Kit — материалы для Ночи кода (локально)

**Не пушить целиком в OnlyGo.** Жюри смотрит историю коммитов.

## Разделение

| Папка | Что | В OnlyGo? |
|-------|-----|-----------|
| `hackathon-kit/` | дизайн-токены, fixture JSON, черновик контракта, snippets | после старта — **по частям** в код (не dump) |
| `../hotel/` | AGENTS.md, skills, промпты агентов | **никогда** |
| `../activity-map-preview/` | визуальный превью стиля | никогда |

## Что здесь

- `design/` — CSS tokens / topology look
- `fixtures/sample-activity.json` — дамп сценария reverse-shell
- `stubs/` — CONTRACT.md, types, bff stub (идеи; реализация пишется в OnlyGo после старта)
- `snippets/tree-layout.ts` — запасной layout; **предпочтительный UI-движок: React Flow + dagre** (см. hotel)
- `PLAN_NOCH_KODA.md` — полный план команды

Правила агентов: **`../hotel/`**.
