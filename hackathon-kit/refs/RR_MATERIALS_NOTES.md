# Выжимка материалов RR + аудит ссылок

Источники (выжимка уже здесь — исходные `.docx` команде не нужны):
- архив PT docx (опционально): `refs/*.docx`
- GitHub: [Runtime-Radar/Runtime-Radar](https://github.com/Runtime-Radar/Runtime-Radar)
- [help](https://github.com/Runtime-Radar/Runtime-Radar/blob/main/docs/guides/help/help.md) · [restapi](https://github.com/Runtime-Radar/Runtime-Radar/blob/main/docs/guides/restapi/restapi.md) · [detectors/guide](https://github.com/Runtime-Radar/Runtime-Radar/blob/main/docs/guides/detectors/guide.md)
- [Tetragon Events](https://tetragon.io/docs/concepts/events/)

## Стек нашего решения

| Слой | Выбор |
|------|--------|
| UI | React + Vite + TS + Zustand + @xyflow/react + dagre + vis-timeline |
| BFF | **Go** `net/http` |
| Данные | public-api / dump JSON (+ опц. webhook) + fixture |

RR сам на Go — не значит писать UI на Go.

## Целевая модель

Канон: `../DATA_MODEL.md`. API и dump → один адаптер. Сеть — наблюдения. Угрозы отдельно. Родитель вне окна не клеится к init.

## Критично для графа

- Рёбра: `process.exec_id` ← `parent.exec_id` / `parent_exec_id` (не pid/ppid).
- Атрибуты: binary, arguments, cwd, uid, `in_init_tree`, start_time, time.
- Сеть: PROCESS_KPROBE → `args[].sock_arg` (не `sockaddr_arg` для tcp_connect/`type: sock`).
- Родитель: `process.parent_exec_id` или `parent.exec_id`.
- Timeline: EXEC+EXIT одного exec_id → range; kprobe → point.
- Bootstrap: отфильтровать cold-start Tetragon по `time`/`start_time`.

## Детектор демо (подтверждено guide.md)

- Function: `inet_csk_listen_start`
- Binary: `*/nc`
- Port: **8888** → CRITICAL, Reason, MITRE **TA0011 / T1571**
- Source UI: **Opening of a socket for incoming connections**
- После `nc -vv -l -p 8888` — отдельный **`ls -al`** (flush в БД)

## Helm / инфра

- Предпочтительно chart **v0.3.0**; admin: `global.administrator.username/password` (не `auth-center.*` из v0.1 Habr).
- Не ставить prometheus/grafana без нужды.
- Mac: Lima / общий Linux; не рассчитывать на голый Docker Desktop + Tetragon.

## Интеграция (3 пути) — важно из МАТЕРИАЛЫ.md

1. **REST public-api (обязателен для сети)** — полный JSON **с `args`**  
   `GET /api/v1/public-api/runtime-event/slice/{left|right}?cursor=<RFC3339>&slice_size=N`  
   заголовок `Access-Token: <personal token>`  
   Docs: [help.md](https://github.com/Runtime-Radar/Runtime-Radar/blob/main/docs/guides/help/help.md), [restapi.md](https://github.com/Runtime-Radar/Runtime-Radar/blob/main/docs/guides/restapi/restapi.md)
2. **ClickHouse / history-api** — быстро и массово, но **без `args`** → **нет IP/путей kprobe**. Не строить Network Map только из CH.
3. **notifier webhook** (опционально) — поток JSONL в реальном времени.

ТЗ «работа с выгрузками» → dump REST (или UI Events export) + желательно Import JSON в UI.

## Sources / объём

- 12 источников (см. help.md): process + **incoming listen** обязательны для демо nc:8888.
- **Не** включать `tcp_sendmsg` на широком ns в режиме All — поток взрывается (десятки тысяч evt/мин).

## Быстрый labeled-корпус без атак

В клоне RR: `event-processor/detector/wasm/*/testdata/{positive,negative}_*.json` (~67 шт.) — точный формат продукта для fixture/тестов нормализатора.

## ML на ночи (Could, не Must)

Offline окна → `scores.json` → BFF merge `risk` на process/window → цвет узлов + подсветка таймлайна + reasons в Incident Summary.  
Не тащить Python в рантайм. Won't: GNN/LoRA/Goat overnight — 55% баллов не про это.

## UX-референсы

Sysdig Inspect (highlights + time brush), Hubble UI (service map), Falco Sidekick (severity colors).

## MoSCoW

Must: 3 панели + sync + fixture + exec_id DAG.  
Should: Attack Path, MITRE, Incident Summary, Import JSON.  
Could: TracingPolicy YAML, Replay, ego-graph, timeline-arrows.
