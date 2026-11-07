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
| `apps/bff/` | **рабочий Go BFF** (fixture/raw/live) | да → OnlyGo `apps/bff/` |
| `scripts/` | `gen-activity.sh`, `dump-rr-events.sh` (Дамир) | да → OnlyGo `scripts/` |
| `design/` | `activity-map.css` + tokens / STYLE_GUIDE | да → `apps/web/src/styles/` |
| `fixtures/` | `sample-activity.json` + `raw/` | да → `apps/bff/fixtures/` / `fixtures/raw/` |
| `stubs/` | web stubs, CONTRACT, types, BASE, CHANGELOG, docs | да → `apps/web`, корень |
| `PLAN_NOCH_KODA.md` | план победы | ориентир; в OnlyGo — README по делу |
| `ARCHITECTURE.md` | вектор продукта | читать всей команде |
| `TEAM_ROLES.md` | Ангелина / Дамир / Денис | задачи по людям |
| `refs/` | PDF / заметки RR | локально / по желанию |

Правила агентов: **`../hotel/`** (не коммитить).
