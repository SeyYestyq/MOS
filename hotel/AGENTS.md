# AGENTS.md (локально / hotel — НЕ коммитить в OnlyGo)

Хакатон: Ночь кода · PT · Activity Map.  
Команда: **3 инженера** — Frontend / Backend / Platform.  
Docs, BASE, питч — общая лёгкая нагрузка в конце, не отдельная роль.

## Стек

- `apps/web`: Vite + React + TS + **Zustand** + **@xyflow/react** + dagre + **vis-timeline**  
- `apps/bff`: **Go** (`net/http`); контракт JSON `/api/activity`  
- Граф процессов: рёбра по **`exec_id` → `parent.exec_id`**, не только pid/ppid  
- Данные: RR public-api / dump **или** fixture  
- Стиль: `../hackathon-kit/design/activity-map.css`  
- Сам RR на Go — не значит писать UI на Go (см. `hackathon-kit/refs/RR_MATERIALS_NOTES.md`)

## Path ownership

| Роль | Можно писать | Нельзя |
|------|--------------|--------|
| **Frontend** | `apps/web/**` | `apps/bff/**`, scripts RR |
| **Backend** | `apps/bff/**`; корень как integrator | `apps/web/**` без joint |
| **Platform** | `scripts/**`, `deploy/**`, `fixtures/raw/**` | `apps/**` без просьбы |

`docs/**`, `CHANGELOG.md`, `BASE.md`, `README.md` — пишут **все** по своим кускам; финальный BASE — joint.

Промпт агенту:  
`Разрешённые пути: … . Остальное read-only. Корень (compose/Makefile) — только Backend-integrator.`

## Контракт / handoff

Контракт = Go structs / JSON schema + `stubs/types.ts`. Breaking change — sync.  
Стык: curl + AC в чате.

## Pacing

scaffold → fixture → tree → panels → sync → summary → live → docs/BASE.

## Skills

| Роль | Файлы |
|------|--------|
| Frontend | `03`, `09` |
| Backend | `02` |
| Platform | `01`, `07` |
| Все | `08`, `11` (ИБ), `05` (freeze), `10` + `04` (docs/питч в конце) |

## DoD

Tree + network + timeline + sync; Incident Summary; fixture; README; CHANGELOG; BASE.md; демо < 60s; Public.
