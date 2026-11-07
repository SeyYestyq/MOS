# Команде (3 человека: A / B / C) — старт после сигнала

Локально: `МосХаб/hotel/AGENTS.md` + skill роли.  
Kit: `МосХаб/hackathon-kit/`. PDF задачи: `hackathon-kit/refs/Задача_Positive_Technologies_описание_задачи.pdf`.

## A — Frontend + Viz
```bash
npm create vite@latest apps/web -- --template react-ts
cd apps/web && npm i @xyflow/react @dagrejs/dagre
mkdir -p src/styles
cp ../../../hackathon-kit/design/{tokens,topology,xyflow-theme,activity-map}.css src/styles/
# main.tsx → import './styles/activity-map.css'
cp ../../../hackathon-kit/stubs/types.ts src/types.ts
```
Дальше: tree → network → timeline → sync → Incident Summary.  
Skill: `03-frontend-activity-map.md`

## B — Data / BFF
```bash
# apps/bff: FastAPI + Pydantic; скопировать fixture + requirements
cp ../hackathon-kit/fixtures/sample-activity.json apps/bff/fixtures/
cp ../hackathon-kit/stubs/bff/requirements.txt apps/bff/
cp ../hackathon-kit/stubs/bff/env.example apps/bff/   # локально → .env
```
Skill: `02-backend-activity-bff.md`. Integrator контракта = B.

## C — Platform + Docs + Pitch
- RR: port-forward `:9000`; Sources; Allow filter ns; All; gen-activity; dump → `fixtures/raw/`
- После старта скопировать в OnlyGo:
```bash
cp -r ../hackathon-kit/stubs/docs docs
cp ../hackathon-kit/stubs/CHANGELOG.md CHANGELOG.md
cp ../hackathon-kit/stubs/BASE.md BASE.md
```
- Вести CHANGELOG после этапов; к freeze заполнить **BASE.md** (блок ИИ).  
Skills: `01-platform-rr.md`, `04-pitch-demo.md`, `10-docs-changelog.md`, `09-soc-demo-language.md`

## Промпт агенту
`prompts/AGENT_BRIEF.md` + путь к skill роли.
