# Старт после сигнала — 3 инженера

## Frontend
```bash
npm create vite@latest apps/web -- --template react-ts
cd apps/web && npm i @xyflow/react @dagrejs/dagre zustand vis-timeline
mkdir -p src/styles
cp ../../../hackathon-kit/design/{tokens,topology,xyflow-theme,activity-map}.css src/styles/
cp ../../../hackathon-kit/stubs/types.ts src/types.ts
```
Skill: `03` (+ `09`). Tree по **exec_id** → network → vis-timeline → Zustand sync → Incident Summary.

## Backend (Go)
```bash
mkdir -p apps/bff/fixtures
cp ../hackathon-kit/stubs/bff/main.go ../hackathon-kit/stubs/bff/go.mod apps/bff/
cp ../hackathon-kit/stubs/bff/env.example apps/bff/
cp ../hackathon-kit/fixtures/sample-activity.json apps/bff/fixtures/
cd apps/bff && go run .
# http://localhost:8000/api/activity
```
Skill: `02`. Normalize exec_id + fixture → live. Integrator корня.

## Platform
- port-forward `:9000`; Sources; Allow filter; All; gen-activity (+ `nc -l 8888`); dump → `fixtures/raw/`  
- Skill: `01`

## Docs / BASE / питч (все, в конце)
```bash
cp -r ../hackathon-kit/stubs/docs docs
cp ../hackathon-kit/stubs/CHANGELOG.md CHANGELOG.md
cp ../hackathon-kit/stubs/BASE.md BASE.md
```
Skills: `10`, `04`, `11`.
