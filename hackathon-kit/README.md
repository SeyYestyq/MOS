# Hackathon Kit — материалы для Ночи кода (локально)

**Не пушить целиком в OnlyGo.** Жюри смотрит историю коммитов.

## Разделение

| Папка | Что | В OnlyGo? |
|-------|-----|-----------|
| `hackathon-kit/` | дизайн-токены, fixture JSON, черновик контракта, snippets | после старта — **по частям** в код (не dump) |
| `../hotel/` | AGENTS.md, skills, промпты агентов | **никогда** |
| `../activity-map-preview/` | визуальный превью стиля | никогда |

## Что здесь

- `design/` — готовый стиль: `activity-map.css` (+ tokens, topology, xyflow-theme, STYLE_GUIDE)
- `fixtures/sample-activity.json` — дамп сценария reverse-shell
- `stubs/` — CONTRACT, types, bff/web stubs (реализация пишется в OnlyGo после старта)
- `snippets/tree-layout.ts` — запасной; **UI: React Flow + dagre** (hotel)
- `PLAN_NOCH_KODA.md` — план команды

Правила агентов: **`../hotel/`**.
