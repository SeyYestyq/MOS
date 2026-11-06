# Команде (4 человека) — шпаргалка старта

## P1 Platform
- `docker info`, `kubectl --context kind-moshub get nodes`
- port-forward RR: `kubectl -n runtime-radar port-forward svc/reverse-proxy 9000:9000`
- экспорт свежих events → дописать `fixtures/` при необходимости

## P2 Backend
- поднять BFF из `stubs/bff/app.stub.py`
- контракт `stubs/CONTRACT.md`
- сначала только fixture, потом RR

## P3 Frontend
- Vite app; подключить `design/tokens.css` + `topology.css`
- реализовать stub `ProcessTreeBoard` + panels
- взять `snippets/tree-layout.ts`

## P4 Pitch
- шаблон презентации уже в OnlyGo
- демо: выбрать suspect python → показать net :4444 → timeline

## Промпт агенту
Вставить целиком `prompts/AGENT_BRIEF.md`
