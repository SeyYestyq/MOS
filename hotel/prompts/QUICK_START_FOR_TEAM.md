# Старт после сигнала — 3 инженера

## 0. Где работаем (обязательно)

1. Клонируете **OnlyGo** с Мос.Хаба — это единственный git для кода.
2. Рядом на диске оставляете `МосХаб/hotel/` (скиллы) и `МосХаб/hackathon-kit/` (заготовки).
3. Из kit **копируете** stubs в OnlyGo и дальше коммитите **только OnlyGo**.
4. Репозиторий `SeyYestyq/MOS` (prep) больше не используете для разработки.

Скиллы остаются в `hotel/` — в OnlyGo их нет.

## Frontend (в клоне OnlyGo)
```bash
npm create vite@latest apps/web -- --template react-ts
cd apps/web && npm i @xyflow/react @dagrejs/dagre zustand vis-timeline
mkdir -p src/styles
# пути к kit поправьте под то, где лежит МосХаб на диске
cp ../../../hackathon-kit/design/{tokens,topology,xyflow-theme,activity-map}.css src/styles/
cp ../../../hackathon-kit/stubs/types.ts src/types.ts
```
Skill: `hotel/skills/03` (+ `09`). Tree по **exec_id** → network → vis-timeline → Zustand sync → Incident Summary.

## Backend (Go, в клоне OnlyGo)
```bash
mkdir -p apps/bff/fixtures
cp ../hackathon-kit/stubs/bff/main.go ../hackathon-kit/stubs/bff/go.mod apps/bff/
cp ../hackathon-kit/stubs/bff/env.example apps/bff/
cp ../hackathon-kit/fixtures/sample-activity.json apps/bff/fixtures/
cd apps/bff && go run .
# http://localhost:8000/api/activity
```
Skill: `hotel/skills/02`. Normalize exec_id + fixture → live. Integrator корня.

## Platform
- port-forward `:9000`; Sources; Allow filter; All; gen-activity (+ `nc -l 8888` + **отдельный `ls -al`**); dump → `fixtures/raw/`  
- Skill: `hotel/skills/01`

## Docs / BASE / питч (все, в конце, в OnlyGo)
```bash
cp -r ../hackathon-kit/stubs/docs docs
cp ../hackathon-kit/stubs/CHANGELOG.md CHANGELOG.md
cp ../hackathon-kit/stubs/BASE.md BASE.md
```
Skills: `10`, `04`, `11`.
