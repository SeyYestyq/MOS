# Старт после сигнала — 3 инженера

## Frontend
```bash
npm create vite@latest apps/web -- --template react-ts
cd apps/web && npm i @xyflow/react @dagrejs/dagre
mkdir -p src/styles
cp ../../../hackathon-kit/design/{tokens,topology,xyflow-theme,activity-map}.css src/styles/
cp ../../../hackathon-kit/stubs/types.ts src/types.ts
```
Skill: `03` (+ `09`). Tree → network → timeline → sync → Incident Summary.

## Backend
```bash
cp ../hackathon-kit/fixtures/sample-activity.json apps/bff/fixtures/
cp ../hackathon-kit/stubs/bff/requirements.txt apps/bff/
cp ../hackathon-kit/stubs/bff/env.example apps/bff/
```
Skill: `02`. Pydantic + `/api/activity` + fixture → live. Integrator корня.

## Platform
- port-forward `:9000`; Sources; Allow filter; All; gen-activity; dump → `fixtures/raw/`  
- Skill: `01`  
- Не сидеть на слайдах — это код/инфра; docs в конце вместе

## Docs / BASE / питч (все, в конце)
```bash
cp -r ../hackathon-kit/stubs/docs docs
cp ../hackathon-kit/stubs/CHANGELOG.md CHANGELOG.md
cp ../hackathon-kit/stubs/BASE.md BASE.md
```
Skills: `10`, `04`. Каждый дописывает CHANGELOG по своим этапам.
