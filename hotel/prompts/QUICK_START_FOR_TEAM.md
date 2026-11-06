# Команде (4 человека) — старт после сигнала

Локально открыть: `МосХаб/hotel/AGENTS.md` + skill своей роли.  
Стили/fixture/stubs: `МосХаб/hackathon-kit/` (в OnlyGo — только куски кода, не AGENTS).

## P1 Platform
- `docker info`, `kubectl --context kind-moshub get nodes`
- `kubectl -n runtime-radar port-forward svc/reverse-proxy 9000:9000`
- при пустых Events — demo-pod / history (см. `.rr-install/ACCESS.md`)
- dump → `fixtures/raw/` для P2

## P2 Backend (первые коммиты)
```bash
# идея со stub — переписать чисто в apps/bff
# модели Pydantic = контракт; GET /api/activity; fixture first
cp ../hackathon-kit/fixtures/sample-activity.json apps/bff/fixtures/
# pytest на схему
```
Skill: `hotel/skills/02-backend-activity-bff.md`

## P3 Frontend (первые минуты — стили уже готовы)
```bash
npm create vite@latest apps/web -- --template react-ts
cd apps/web && npm i @xyflow/react @dagrejs/dagre
mkdir -p src/styles
cp ../../../hackathon-kit/design/{tokens,topology,xyflow-theme,activity-map}.css src/styles/
# main.tsx → import './styles/activity-map.css'
cp ../../../hackathon-kit/stubs/types.ts src/types.ts
```
Дальше: React Flow дерево → network → timeline → sync → Incident Summary.  
Skill: `hotel/skills/03-frontend-activity-map.md`  
Глаза на стиль: `npm run dev` в `activity-map-preview` (не копировать оптом).

## P4 Pitch
- шаблон pptx уже в OnlyGo
- демо: suspect → net :4444 → timeline → copy Incident Summary  
Skill: `hotel/skills/04-pitch-demo.md`

## Промпт агенту
Вставить `prompts/AGENT_BRIEF.md` + путь к своему `skills/0N-….md`.
