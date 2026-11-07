# Выжимка материалов RR + аудит ссылок

Источники:
- `../Материалы по Runtime Radar.docx`
- `../Материалы для карты активности контейнера.docx`
- GitHub: [Runtime-Radar/runtime-radar](https://github.com/Runtime-Radar/runtime-radar)
- [detectors/guide.md](https://github.com/Runtime-Radar/runtime-radar/blob/main/docs/guides/detectors/guide.md)
- [Tetragon Events](https://tetragon.io/docs/concepts/events/)

## Стек нашего решения

| Слой | Выбор |
|------|--------|
| UI | React + Vite + TS + Zustand + @xyflow/react + dagre + vis-timeline |
| BFF | **Go** `net/http` |
| Данные | public-api / dump JSON (+ опц. webhook) + fixture |

RR сам на Go — не значит писать UI на Go.

## Критично для графа

- Рёбра: `process.exec_id` ← `parent.exec_id` / `parent_exec_id` (не pid/ppid).
- Атрибуты: binary, arguments, cwd, uid, `in_init_tree`, start_time, time.
- Сеть: PROCESS_KPROBE → `args[0].sock_arg` (saddr/sport/daddr/dport/family).
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

## Интеграция (3 пути)

1. public-api через reverse-proxy (:32000 / :9000) + **JSON выгрузка Events**
2. history-api / ClickHouse (port-forward)
3. notifier webhook (опционально)

ТЗ «работа с выгрузками» → dump + желательно Import JSON в UI.

## UX-референсы

Sysdig Inspect (highlights + time brush), Hubble UI (service map), Falco Sidekick (severity colors).

## MoSCoW

Must: 3 панели + sync + fixture + exec_id DAG.  
Should: Attack Path, MITRE, Incident Summary, Import JSON.  
Could: TracingPolicy YAML, Replay, ego-graph, timeline-arrows.
