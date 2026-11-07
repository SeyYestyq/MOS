# Старт после сигнала — Ангелина / Дамир / Денис

Полные ветки: `hackathon-kit/TEAM_ROLES.md`.

## 0. Где работаем

1. Клон **OnlyGo** (Мос.Хаб) — единственный git для кода.
2. Рядом: `МосХаб/hotel/` (скиллы) + `МосХаб/hackathon-kit/` (заготовки).
3. `cp` stubs в OnlyGo → коммиты только в OnlyGo. MOS/prep не для фич.

---

## Ангелина — Frontend
```bash
npm create vite@latest apps/web -- --template react-ts
cd apps/web && npm i @xyflow/react @dagrejs/dagre zustand vis-timeline
mkdir -p src/styles
cp ../../../hackathon-kit/design/{tokens,topology,xyflow-theme,activity-map}.css src/styles/
cp ../../../hackathon-kit/stubs/types.ts src/types.ts
```
Skills: `03`, `09`. Порядок: tree → network → timeline → Zustand → Summary.  
Можно начинать **сразу на fixture**, не ждать live RR.

---

## Денис — Backend
```bash
mkdir -p apps/bff/fixtures
cp ../hackathon-kit/stubs/bff/main.go ../hackathon-kit/stubs/bff/go.mod apps/bff/
cp ../hackathon-kit/stubs/bff/env.example apps/bff/
cp ../hackathon-kit/fixtures/sample-activity.json apps/bff/fixtures/
cd apps/bff && go run .
# http://localhost:8000/api/activity
```
Skills: `02`, `08`, `11`. Нормализатор exec_id + sock_arg; `go test`; XSS; сверка curl↔UI.  
Raw dump ждёт от Дамира → подключить в нормализатор.

---

## Дамир — DevOps + live-стык
- Linux/Lima; Helm **v0.3.0**; Sources process + **incoming listen**; Allow `hackathon`; All.
- `gen-activity.sh` (+ **ls после nc**); REST slice dump → `fixtures/raw/`; curl + Access-Token Денису.
- Skills: `01`, `07`. Live wiring BFF jointly с Денисом.

---

## Docs / BASE (все, в конце)
```bash
cp -r ../hackathon-kit/stubs/docs docs
cp ../hackathon-kit/stubs/CHANGELOG.md CHANGELOG.md
cp ../hackathon-kit/stubs/BASE.md BASE.md
```
Skills: `10`, `04`, `11`.
