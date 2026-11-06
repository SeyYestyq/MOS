# AGENT BRIEF — Ночь кода / Интерактивная карта активности контейнера

Вставлять агенту **после старта**. Файлы из `МосХаб/hotel/` и `hackathon-kit/` **не коммитить** в OnlyGo как «правила для ИИ».

## Задача

Веб-интерфейс по данным Runtime Radar в **одном окне**:
1. хронология действий,
2. дерево процессов (ppid→pid),
3. сетевые соединения,
4. аналитик быстро восстанавливает картину внутри контейнера.

## Критерии жюри

30% решение · 25% качество · 20% инновация · 15% презентация · 10% потенциал

## Локальный контекст (диск, не git сдачи)

- Kind `moshub`, RR ns `runtime-radar`, UI `:9000` admin/Password
- `МосХаб/hotel/AGENTS.md` + `hotel/skills/0N-….md` по роли — **единственные** инструкции; `claude-skills-main` не читать
- `МосХаб/hackathon-kit/design/` — стиль
- `МосХаб/hackathon-kit/fixtures/sample-activity.json` — fixture атаки
- Preview стиля: `activity-map-preview` (глаза, не копировать оптом в OnlyGo)

## Стек (зафиксировано)

- Web: Vite + React + TS + **`@xyflow/react` + dagre** (не invent SVG graph math)
- BFF: FastAPI + **Pydantic** (= OpenAPI); pytest на схему
- Fixture-first: `X-Data-Source: fixture|live`
- Live RR: HTTP через reverse-proxy / public-api (см. `.rr-install/ACCESS.md`), не парсинг логов пода

## Зона агента

В промпте всегда: разрешённые пути только. Корень репо — только integrator.

## Инновация (заложить в UI)

Помимо Replay: блок **Incident Summary** — текст для SOC по выбранному процессу + copy  
(пример: pod, pid chain, external IP:port, время).

## Pacing коммитов в OnlyGo

Дробить этапы. Не один mega-commit через 15 минут после старта.  
В OnlyGo: код, OpenAPI/модели, README, презентация — **без** AGENTS.md/skills/.

## DoD

Tree + network + timeline + sync + Incident Summary; fixture fallback; README; демо time-to-insight < 60s.
