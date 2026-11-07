# AGENT BRIEF — Activity Map (PT / Runtime Radar)

После старта работаем **только в OnlyGo** (Мос.Хаб).  
`hotel/` (скиллы) — отдельная папка на диске, в OnlyGo **не** коммитить.  
Stubs из `hackathon-kit/` — перенести в OnlyGo и реализовывать план там; prep/MOS не продолжать.

## Задача

Одно окно: хронология + дерево процессов + сеть — аналитик ИБ восстанавливает картину быстро.

## Роли (3 инженера)

| Роль | Зона | Skill |
|------|------|-------|
| **Frontend** | `apps/web/**` | `03` (+ `09`) |
| **Backend** | `apps/bff/**` | `02` |
| **Platform** | `scripts/**`, `fixtures/raw/**` | `01` |

Docs / `BASE.md` / питч — **общая** работа в конце, не вместо кода.

## Стек

Vite/React/xyflow+dagre+vis-timeline+Zustand · **Go BFF** · fixture-first · live RR по HTTP.

## Сдача

Код + docs + **BASE.md** (ИИ) + презентация + демо < 60s.

## DoD

Tree + network + timeline + sync + Incident Summary; fixture; README; CHANGELOG; BASE; Public; без hotel в git.
