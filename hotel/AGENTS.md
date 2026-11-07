# AGENTS.md (локально / hotel — НЕ коммитить в OnlyGo)

Хакатон: Ночь кода · PT · Activity Map.  
Команда: **Ангелина** (Frontend) · **Дамир** (DevOps + Backend-интегратор) · **Денис** (Backend).  
Ветки задач: `../hackathon-kit/TEAM_ROLES.md`. Docs/BASE/питч — jointly в конце.

## Где код

- **Писать и пушить:** только репозиторий **OnlyGo** на Мос.Хабе.
- **Скиллы / этот файл:** папка `hotel/` рядом на диске — отдельно, не в git сдачи.
- **Заготовки:** `hackathon-kit/` → **перенести** в OnlyGo (`cp` stubs/design/fixtures), дальше жить в OnlyGo.
- Репозиторий подготовки `MOS` / ветка `prep` — не место разработки продукта.

## Стек

- `apps/web`: Vite + React + TS + **Zustand** + **@xyflow/react** + dagre + **vis-timeline**  
- `apps/bff`: **Go** (`net/http`); контракт JSON `/api/activity`  
- Граф процессов: рёбра по **`exec_id` → `parent.exec_id`**, не только pid/ppid  
- **Данные:** `hackathon-kit/DATA_MODEL.md` (канон снимка)  
- Стиль: `../hackathon-kit/design/activity-map.css`

## Path ownership

| Кто | Можно писать | Нельзя |
|-----|--------------|--------|
| **Ангелина** | `apps/web/**` | `apps/bff/**`, scripts RR |
| **Денис** | `apps/bff/**` (+ tests); контракт | `apps/web/**` без joint; не ломать RR-стенд Дамира |
| **Дамир** | `scripts/**`, `deploy/**`, `fixtures/raw/**`, env.example, compose; live-стык BFF jointly | UI Ангелины без просьбы; нормализатор без согласования с Денисом |

`docs/**`, `CHANGELOG.md`, `BASE.md`, `README.md` — все по кускам; финальный BASE — joint.

Промпт агенту:  
`Разрешённые пути: … . Остальное read-only.`

## Контракт / handoff

Контракт = Go structs + `stubs/types.ts` — владелец **Денис**. Breaking change — sync с Ангелиной.  
Raw dump + curl slice — **Дамир** → Денис. Стык: curl + AC в чате.

## Skills

| Кто | Файлы |
|-----|--------|
| Ангелина | `03`, `09` |
| Денис | `02`, `08`, `11` |
| Дамир | `01`, `07`, куски `02` (live) |
| Все | `05` (freeze), `10` + `04` (docs/питч) |

## DoD

Tree + network + timeline + sync; Incident Summary; fixture; README; CHANGELOG; BASE.md; демо < 60s; Public; `go test` зелёный.
