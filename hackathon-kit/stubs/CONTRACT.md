# API contract (зафиксировать в час 0)

## `GET /api/activity?pod=<name>&namespace=<ns>`

Ответ соответствует типам из `hackathon-kit/stubs/types.ts` (и форме `fixtures/sample-activity.json`).

### Fields

- `processes[]`:
  - `id`: string (равен `execId`)
  - `execId`: string (уникальный 64-битный Tetragon hash в base64)
  - `parentExecId`: string | null (ссылка на родительский `exec_id`)
  - `pid`: number
  - `ppid`: number (опционально, для совместимости)
  - `binary`: string (путь к исполняемому файлу)
  - `arguments`: string (аргументы командной строки)
  - `cwd`: string (рабочая директория)
  - `role`: "entry" | "shell" | "tool" | "suspect" | string
  - `inInitTree`: boolean (отделяет фоновый init-процесс контейнера от `kubectl exec`)
  - `startTime`: string ISO8601
  - `firstSeen`: string ISO8601
  - `lastSeen`: string ISO8601
  - `mitre`: string[] (например `["TA0011", "T1571"]`)

- `network[]`:
  - `id`: string
  - `execId`: string (привязка сокета к процессу, открывшему соединение)
  - `fromPid`: number (legacy)
  - `saddr`: string
  - `sport`: number
  - `daddr`: string
  - `dport`: number
  - `family`: string (AF_INET / AF_INET6)
  - `dst`: string (целевой хост или DNS имя)
  - `port`: number (порт назначения)
  - `proto`: string ("tcp" | "udp")
  - `label`: string ("listen" | "egress" | "download")
  - `t`: string ISO8601

- `events[]`:
  - `id`: string
  - `t`: string ISO8601
  - `type`: string ("PROCESS_EXEC" | "PROCESS_EXIT" | "PROCESS_KPROBE")
  - `execId`: string
  - `pid`: number
  - `summary`: string
  - `reason`: string (контекст Wasm-детектора, например "Listening on port 8888")

### Tree rule (Математическая модель DAG)

1. **Основное правило:** Ребро `parent → child` строится **строго по `child.parentExecId === parent.execId`**.
2. **Корень (Root):** Узлы, у которых `parentExecId === null` или родитель отсутствует в текущей выборке (init-процесс контейнера или корень сессии `kubectl exec`).
3. **Защита от PID reuse:** Запрещено строить дерево только по `ppid === pid`, так как в контейнерах и при агрегации нескольких подов PID не уникальны. `ppid` допускается только как запасной fallback для моков.

### Fallback

Если Runtime Radar API недоступен → вернуть `fixtures/sample-activity.json` со статусом HTTP 200 и заголовком `X-Data-Source: fixture`. При живом ответе — заголовок `X-Data-Source: live`.
