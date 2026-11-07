# Три ветки команды OnlyGo

Синхронизация каждые **45–60 мин**. Контракт API не ломать без чата.  
Код только в **OnlyGo** (Мос.Хаб). Скиллы — `hotel/` на диске. Вектор — `ARCHITECTURE.md`.

| Кто | Ветка | Зона git | Skills |
|-----|-------|----------|--------|
| **Ангелина** | Frontend | `apps/web/**` | `03`, `09` |
| **Дамир** | DevOps + Backend (интегратор) | `scripts/**`, `deploy/**`, `fixtures/raw/**`, корень compose/Makefile; стык live BFF | `01`, `07`, `02` (live/интеграция) |
| **Денис** | Backend | `apps/bff/**`, `apps/bff/**/*_test.go`, контракт | `02`, `08`, `11` |

Docs / `BASE.md` / питч — все в конце. Когда Ангелина закрывает Must по UI — подключается к docs/питчу и полировке sync.

---

## Ветка A — Ангелина (Frontend)

**Цель:** Single Pane of Glass: дерево + сеть + таймлайн + sync. UI живёт на fixture, не ждёт кластер.

**Старт сейчас:** scaffold из `hackathon-kit` → дерево по `execId` → сеть → timeline → Zustand → Summary.

| ID | Задача | Pri |
|----|--------|-----|
| A1 | Vite React TS + `activity-map.css` + `types.ts` | Must |
| A2 | Layout: Highlights + 3 панели + drawer | Must |
| A3 | Process Tree: dagre; ребро только если родитель в снимке; `outside_window` явно | Must |
| A4 | Network: наблюдения `connect/listen/close`, не «живые сокеты» | Must |
| A5 | Timeline: жизнь процесса по started/exited; 3-я дорожка = `threats[]` | Must |
| A6 | Zustand: `selectedExecId` + time brush на все панели | Must |
| A7 | Incident Summary + copy; MITRE badges | Should |
| A8 | Import JSON dump; Attack Path highlight | Should |
| A9 | TracingPolicy YAML / Replay / ego-graph | Could |

**DoD Ангелины:** на fixture видно дерево+сеть+таймлайн, клик синхронизирует панели, демо-цепочка nc:8888 читается без Backend live.

**Handoff от Дениса:** стабильный `GET /api/activity` + `sample-activity.json` / proxy.

---

## Ветка B — Дамир (DevOps + Backend-интегратор)

**Цель:** Runtime Radar живой, сценарий атаки воспроизводим, dump с `args` у Дениса, live-стык BFF↔RR, корень репо не разваливается.

| ID | Задача | Pri |
|----|--------|-----|
| B1 | Linux-стенд (не голый Docker Desktop Mac); BTF; Helm **v0.3.0** + `global.administrator.*` | Must |
| B2 | Sources: process + **incoming listen**; Allow `hackathon`; **All**; без широкого `tcp_sendmsg` | Must |
| B3 | `scripts/gen-activity.sh` (4 фазы + **ls после nc**) | Must |
| B4 | Dump **REST slice** → `fixtures/raw/events.json` (не CH для сети); запас — testdata детекторов | Must |
| B5 | Таблица URL/auth + рабочий `curl` + `Access-Token` для Дениса (до T4) | Must |
| B6 | Port-forward / NodePort; `.env.example` без секретов в git | Must |
| B7 | Live-режим BFF: проводка public-api → нормализатор Дениса; `X-Data-Source: live` | Should |
| B8 | compose/Makefile / README «как поднять стенд» | Should |
| B9 | 10-мин triage RR; если eBPF мёртв → не хоронить ночь, усилить fixture | Must-дисциплина |

**DoD Дамира:** Events в RR с threat на nc:8888; raw dump у команды; curl slice работает; при живом RR UI может переключиться на live.

**Стык с Денисом:** Дамир даёт raw JSON + curl; Денис нормализует. Live-флаг / env — jointly, ломает контракт только по согласованию.

---

## Ветка C — Денис (Backend)

**Цель:** Go BFF — единственный контракт для фронта; строгий DAG; сеть из `sock_arg`; тесты; XSS; fixture-first.

| ID | Задача | Pri |
|----|--------|-----|
| C1 | Scaffold `apps/bff`: `/healthz`, `/api/activity`, CORS, `X-Data-Source` | Must |
| C2 | **Один контракт** по `DATA_MODEL.md`: Go structs = JSON = TS; тест фикстуры ловит расхождение | Must |
| C3 | **Один адаптер** API и dump → одинаковые события; честный parent; времена; threats[]; net.kind | Must |
| C4 | Fixture `sample-activity.json` + разбор `fixtures/raw/` от Дамира | Must |
| C5 | XSS sanitize arguments; **нет** `dangerously…` на бэке в HTML | Must |
| C6 | `go test ./...`: DAG без pid-only рёбер; fixture roundtrip; sanitize | Must |
| C7 | Чеклист «всё сходится»: curl activity → поля tree/net/events; сверка с UI Ангелины | Must |
| C8 | Smart Noise Folding | Should |
| C9 | Merge optional `risk` / `/api/windows` / `/api/insight` | Could |

**DoD Дениса:** `go run` отдаёт activity на fixture; тесты зелёные; после dump Дамира — нормализация без ручной правки JSON в UI; контракт совпадает с `types.ts`.

**Стык с Ангелиной:** один JSON; breaking change только после ping. Proxy Vite → `:8000`.

---

## Порядок ночи (кто кого ждёт)

```
T0–T1  Ангелина: scaffold+css     Денис: BFF+fixture     Дамир: Helm/RR/Sources
T1–T2  Ангелина: tree             Денис: нормализатор     Дамир: gen-activity + dump
T2–T3  Ангелина: net+timeline     Денис: raw→activity     Дамир: curl slice / Access-Token
T3–T4  Ангелина: Zustand sync     Денис: tests+XSS        Дамир: live wiring
T4–T5  Ангелина: Summary/MITRE    Денис: Noise Folding?   Дамир: стабильность стенда
T5–T7  joint: E2E, BASE, питч, freeze
```

Ангелина после закрытия Must UI → docs/питч + полировка.  
Денис отвечает за «бэкенд работает и доказан тестами».  
Дамир отвечает за «данные из мира доходят».
