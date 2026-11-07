# Hackathon Kit — подготовка к Ночи кода

## Где работает команда

| Место | Роль |
|-------|------|
| **OnlyGo** (`git@hub.mos.ru:event_DIT/2026/onlygo.git`) | **единственный** репозиторий разработки и сдачи |
| **этот MOS / prep** | бэкап плана и заготовок; **не** писать сюда фичи продукта |
| **`../hotel/`** | AGENTS + skills — **отдельная папка на диске**, в OnlyGo **никогда** |

После старта: клон OnlyGo → `cp` stubs/design/fixtures из kit → реализация плана в OnlyGo.  
Подробно: `DO_NOT_PUSH.txt`, `../hotel/prompts/QUICK_START_FOR_TEAM.md`.

## Что здесь (источник для переноса)

| Папка | Что | В OnlyGo? |
|-------|-----|-----------|
| `design/` | `activity-map.css` + tokens / STYLE_GUIDE | да → `apps/web/src/styles/` |
| `fixtures/` | `sample-activity.json` | да → `apps/bff/fixtures/` |
| `stubs/` | Go BFF, web stubs, CONTRACT, types, BASE, CHANGELOG, docs | да → `apps/*`, корень |
| `PLAN_NOCH_KODA.md` | план победы | ориентир; в OnlyGo — README по делу |
| `ARCHITECTURE.md` | вектор продукта: слои, модель данных, UX-цель | читать всей команде |
| `refs/` | PDF / заметки по материалам RR | локально / по желанию |

Правила агентов: **`../hotel/`** (не коммитить).
