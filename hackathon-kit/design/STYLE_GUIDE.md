# Style guide — Activity Map

Единственный источник стиля: **`hackathon-kit/design/`**  
(не искать токены в чужих репозиториях).

| Файл | Назначение |
|------|------------|
| `tokens.css` | CSS variables, body, Manrope |
| `topology.css` | панели `.admin-surface`, рамки, canvas/legend |
| `xyflow-theme.css` | тёмная тема React Flow + кастомные process-ноды |
| `activity-map.css` | **один импорт** = всё выше |

Живой превью (глаза): `../activity-map-preview` — после старта **не** копировать целиком в OnlyGo; копировать CSS из `design/`.

## После старта (P3, 1-я минута)

```bash
# из корня OnlyGo после scaffold apps/web
cp ../hackathon-kit/design/tokens.css apps/web/src/styles/
cp ../hackathon-kit/design/topology.css apps/web/src/styles/
cp ../hackathon-kit/design/xyflow-theme.css apps/web/src/styles/
cp ../hackathon-kit/design/activity-map.css apps/web/src/styles/
# в main.tsx: import './styles/activity-map.css'
```

Типы/fixture рядом:

```bash
cp ../hackathon-kit/stubs/types.ts apps/web/src/
cp ../hackathon-kit/fixtures/sample-activity.json apps/web/public/  # или через BFF
```

## Визуал

- Dark body gradient + Manrope
- Карточки `.admin-surface` + `.liquid-metal-frame` / `--soft`
- Узлы процессов: цвет stroke по `role` (см. таблицу)
- Сетка / React Flow Background gap ≈ 28px
- Legend снизу

## Цвета ролей процессов

| role | color | смысл |
|------|-------|--------|
| entry | `#37d5ff` | entrypoint |
| shell | `#ffce63` | shell |
| tool | `#84cc16` | nginx/curl/… |
| suspect | `#ff5f8d` | payload / reverse-shell / miner |

CSS vars: `--role-entry`, `--role-shell`, `--role-tool`, `--role-suspect` в `tokens.css`.

## Layout MVP

```
[ Process tree (React Flow) ..... ] [ Selected / Incident Summary ]
[                                ] [ Network  ]
[                                ] [ Timeline ]
```

## Не тащить в OnlyGo

- чужой бизнес-домен / брендинг
- весь `activity-map-preview` как готовое «решение одним коммитом»
