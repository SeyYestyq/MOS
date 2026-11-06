# Style guide (из admin new4)

Источник визуала:
- `/Users/damir/Documents/NEW/11.07.26_new/new4/frontend/src/index.css` (токены)
- `.../liquid-metal.css` (рамки)
- `.../App.css` (`.admin-surface`, `.admin-topology-*`)
- Компоненты-референс: `ServiceTopology.tsx`, `PipelineFlowMap.tsx`

## Что копируем

- Dark body gradient + Manrope
- Карточки `.admin-surface` + `.liquid-metal-frame--soft`
- SVG-узлы: rounded rect 12px, цветной stroke по статусу/роли, selection ring
- Сетка на canvas 28px
- Legend dots снизу

## Что НЕ копируем

- AdminPanel целиком, ServicesPanel
- Бизнес-слои core/middleware/user-facing (для хакатона — дерево процессов L0/L1/L2)
- Чужой продуктовый брендинг / доменные лейблы

## Цвета ролей процессов (предложение)

| role | color | смысл |
|------|-------|--------|
| entry | `#37d5ff` | entrypoint |
| shell | `#ffce63` | shell |
| tool | `#84cc16` | nginx/curl/… |
| suspect | `#ff5f8d` | payload / reverse-shell / miner |

## Layout MVP

```
[ Process tree SVG .............. ] [ Selected ]
[                                ] [ Network  ]
[                                ] [ Timeline ]
```
