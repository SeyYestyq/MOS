# Старт после сигнала — Ангелина / Дамир / Денис

Полные ветки: `hackathon-kit/TEAM_ROLES.md`.

## 0. Где работаем

1. Рабочий репозиторий: **`МосХаб/OnlyGo`** → `git@hub.mos.ru:event_DIT/2026/onlygo.git`.
2. Рядом: `МосХаб/hotel/` (скиллы) + `МосХаб/hackathon-kit/` (справочник).
3. Пишем и пушим **только в OnlyGo**. MOS/prep — не для фич продукта.
4. BFF уже в `OnlyGo/apps/bff` (Go). Web — `apps/web` (Ангелина).

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

## Денис — Backend (Go)
```bash
cp -r ../hackathon-kit/apps/bff apps/bff
cd apps/bff && go test ./... && go run .
# http://localhost:8000/api/activity  → X-Data-Source: fixture
```
Skills: `02`, `08`, `11`. Дожимать нормализатор под реальный dump; `go test`; сверка curl↔UI.  
Уже есть: sanitize, fixture, raw normalize, live client skeleton.

---

## Дамир — DevOps + live-стык
```bash
cp -r ../hackathon-kit/scripts scripts
# Sources process + incoming listen; Allow hackathon; All
./scripts/gen-activity.sh
RR_TOKEN=... RR_BASE_URL=https://localhost:9000 ./scripts/dump-rr-events.sh
```
Skills: `01`, `07`. Live: `DATA_MODE=live` в `apps/bff` jointly с Денисом.

---

## Docs / BASE (все, в конце)
```bash
cp -r ../hackathon-kit/stubs/docs docs
cp ../hackathon-kit/stubs/CHANGELOG.md CHANGELOG.md
cp ../hackathon-kit/stubs/BASE.md BASE.md
```
Skills: `10`, `04`, `11`.
