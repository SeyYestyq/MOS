# Генеральный план победы команды OnlyGo — «Ночь кода 2026»

**Задача:** №2 «Интерактивная карта активности контейнера»  
**Партнёр:** Positive Technologies (команда PT Container Security & Runtime Radar)  
**Команда:** **3 человека** (Frontend Engineer / Backend Engineer / Platform Engineer)  
**Репозиторий сдачи:** `OnlyGo` на платформе «Мос.Хаб» (`git@hub.mos.ru:event_DIT/2026/onlygo.git`)  
**Техническая и нормативная база:**
- Описание задачи: `OnlyGo/Задача_Positive_Technologies_описание_задачи.pdf`
- Критерии оценки: `OnlyGo/Критерии_оценки_задачи_Интерактивная_карта_активности_контейнера.pdf`
- Требования к ПО: `OnlyGo/Требования_к_ПО_задачи_Интерактивная_карта_активности_контейнера.pdf`
- Стенограмма брифинга эксперта PT (Святослав Сорокин): `Новая запись 4.txt`
- Архитектурный гид и референсы: `Материалы по Runtime Radar.docx`
- Реестр технических материалов и отраслевых UX-паттернов: `Материалы для карты активности контейнера.docx`

### 0. Где работаем командой (обязательно)

| Что | Где |
|-----|-----|
| Весь код, коммиты, сдача жюри | **OnlyGo** на Мос.Хабе |
| План + stubs/design/fixtures (заготовки) | локально `hackathon-kit/` → **перенести** в OnlyGo после старта |
| AGENTS.md + skills агентов | **`hotel/`** — отдельная папка на диске, **не** в git OnlyGo |
| `github.com/SeyYestyq/MOS` (ветка `prep`) | только бэкап подготовки; **не** основной реп разработки |

После сигнала: clone OnlyGo → `cp` из kit по `hotel/prompts/QUICK_START_FOR_TEAM.md` → реализовать этот план **в OnlyGo**. Skills читать из `hotel/skills/`.

---

## 1. Понимание «боли» экспертов PT и продуктовый контекст

### 1.1. В чём главная боль аналитика ИБ и руководства
*(из брифинга старшего разработчика PT CS и Runtime Radar Святослава Сорокина)*
> *«Подходит руководство к безопасному: "А что случилось? Ну-ка расскажи!" И чтобы достаточно оперативно это сделать и не нервировать никого, кто это спрашивает, нам, в идеале, сделать удобную картинку (UI/UX), в которой будет видно: порядок действий (таймлайн), родственные связи (дерево процессов), события и сетевое взаимодействие. Не идти по какому-то дикому бешеному списку, особенно если ловились ВСЕ события (режим ALL) — и негативные, и позитивные, а сразу безопасному проверить, посмотреть, понять и осознать за считанные секунды».*

В ситуации реального инцидента специалист должен ответить руководству на три ключевых вопроса:
1. **Что произошло в контейнере?** (какие команды запускались, какие файлы читались, какие порты открывались).
2. **Как развивалась цепочка атаки?** (кто кого породил: от внешней инъекции `kubectl exec` через шелл к вредоносному пейлоаду).
3. **Почему сработали детекторы и как закрыть угрозу?** (какая сигнатура ядра сработала — например, `inet_csk_listen_start` — и как за 1 клик сгенерировать правило блокировки).

### 1.2. Три системных тупика существующих инструментов
1. **Когнитивный коллапс в режиме `All`:** Полное логирование генерирует тысячи вызовов ядра (`ls`, `cat`, healthcheck-пробы). В плоской таблице атака моментально тонет в шуме.
2. **Фатальная ошибка связывания процессов по `PID/PPID`:** В контейнерах PID быстро переиспользуются ядром, а при агрегации нескольких подов происходят коллизии. **Строгое требование PT:** строить граф как **строгий DAG по `process.exec_id` $\rightarrow$ `parent.exec_id`**, генерируемым eBPF-сенсором Cilium Tetragon (уникальный 64-битный base64-хеш на уровне кластера).
3. **Разорванность контекста и хаос раскладки:** Сетевые соединения живут отдельно от процессов, а таймлайн оторван от дерева. Использование физических симуляций (вроде `d3-force`) превращает дерево процессов в «спагетти» — для дерева процессов требуется строгая иерархическая укладка **Dagre (Top-Bottom / Left-Right)**.

### 1.3. Наш ответ: 15-секундный «Time-to-Insight» в концепции Single Pane of Glass
Единый экран расследования (Single Pane of Glass), объединяющий 4 согласованных квадранта:
- **Instant Highlights (Sysdig Inspect pattern):** карточки-индикаторы с мгновенным фильтром по угрозам.
- **Process Tree Canvas (`@xyflow/react` + Dagre):** иерархическое дерево с поддержкой паттерна `ego-graph` и отделением фонового `in_init_tree` от интерактивных сессий `kubectl exec`.
- **Network Map (Hubble UI style):** топология сервисов и сокетов, привязанных к конкретным `exec_id`.
- **3-Lane Synchronized Timeline (`vis-timeline` + `timeline-arrows`):** многодорожечная временная шкала с интервалами жизни, сетевыми kprobe-вызовами и стрелками порождения процессов.
- **Context & Response Drawer:** карточка инцидента с генерацией отчета для CISO и **экспортом YAML-политики TracingPolicy** для мгновенной блокировки повторной атаки.

---

## 2. Критерии жюри и стратегия получения 10/10

| Критерий | Вес | Ожидание экспертов Positive Technologies | Наше инженерное решение для получения 10/10 |
|---|:---:|---|---|
| **Решение конкурсного задания** | **30%** | Полное объединение 3 видов представлений (дерево, сеть, таймлайн) в едином интерфейсе на реальных данных `hackathon`. | **Single Pane of Glass (4 квадранта):** Process Tree + Network Topology + 3-Lane `vis-timeline` в общем окне без перезагрузки и переключения вкладок. Реальные данные из k8s через Go BFF. |
| **Качество решения** | **25%** | Корректный DAG без коллизий PID; отсутствие хаоса в режиме `All`; чистота архитектуры; DevSecOps (защита от XSS). | Строгое построение DAG по `exec_id` / `parent_exec_id`; умное схлопывание шума (**Noise Folding**); **Anti-XSS защита** (экранирование `process.arguments`); двухрежимный Go BFF с fallback (`X-Data-Source: live \| fixture`). |
| **Инновационность и креативность** | **20%** | Передовые UX-инструменты расследования для аналитика SOC, вызывающие вау-эффект. | **1. Brushing & Linking** (Zustand v5);<br>**2. Sub-second Time-Window Brushing** (выделение диапазона на таймлайне отсекает шум верхних графов);<br>**3. Attack Path Highlight** с бейджами MITRE ATT&CK (TA0011, T1571);<br>**4. 1-Click TracingPolicy Generator** (создание YAML политики Tetragon для блокировки угрозы). |
| **Презентация и командная работа** | **15%** | Четкий питч, демонстрация ценности, слаженность ролей. | Сюжетный сторителлинг **«15 секунд расследования инцидента»**: от ночного звонка директора до копирования отчета и блокировки. Слаженная защита 3 инженеров по таймингу (3 минуты). |
| **Потенциал использования** | **10%** | Готовность к внедрению в open-source экосистему Runtime Radar. | Архитектура встраиваемого веб-модуля, стилистически повторяющего дизайн-систему `radar-ui`, работающего с `public-api`, `history-api` (ClickHouse) и Webhook-уведомлениями `notifier`. |

---

## 3. Архитектура и математическая модель корреляции данных

### 3.1. Полный цикл движения телеметрии eBPF $\rightarrow$ UI

```
  ┌────────────────────────────────────────────────────────────────────────┐
  │                           Kubernetes Кластер                           │
  │                                                                        │
  │  [ Pod: attacker-pod ] (вызовы: execve, socket, connect, listen)       │
  │          │                                                             │
  │          ▼                                                             │
  │  [ runtime-monitor (Cilium Tetragon eBPF DaemonSet) ]                  │
  │          │  обогащение: namespace, pod, start_time, exec_id, sock_arg  │
  │          ▼                                                             │
  │  [ event-processor (Go Wasm TinyGo 0.41.1 детектор угроз) ]            │
  │          │  маркировка: inet_csk_listen_start -> TA0011 / T1571       │
  │          ▼                                                             │
  │  [ ClickHouse (history-api: 8003/9003) / PostgreSQL / RabbitMQ ]       │
  │          │                                                             │
  │  [ public-api / reverse-proxy (NodePort :32000 или port-forward :9000)]│
  │  [ notifier (опциональный Webhook стриминг через Go-шаблоны) ]         │
  └──────────┬─────────────────────────────────────────────────────────────┘
             │ HTTP REST / Events Dump JSON / Webhook
             ▼
  ┌────────────────────────────────────────────────────────────────────────┐
  │                       OnlyGo Backend (Go BFF :8000)                    │
  │  - Нормализатор: плоский поток eBPF -> Строгий DAG (exec_id -> parent) │
  │  - Экстрактор сокетов: sock_arg (saddr, sport, daddr, dport) -> net    │
  │  - Smart Noise Folding: схлопывание фоновых вызовов в режиме All       │
  │  - Fallback engine: live RR -> raw fixture (X-Data-Source)             │
  │  - Security Sanitizer: нейтрализация XSS-пейлоадов в cmdline           │
  └──────────┬─────────────────────────────────────────────────────────────┘
             │ GET /api/activity?namespace=hackathon&pod=attacker-pod
             ▼
  ┌────────────────────────────────────────────────────────────────────────┐
  │                     OnlyGo Frontend (React + Vite :5173)               │
  │  ┌──────────────────────────────────────────────────────────────────┐  │
  │  │ Header: Pod Selector + Instant Highlights (3 KPI Cards)           │  │
  │  ├──────────────────────────────────┬───────────────────────────────┤  │
  │  │ Process Tree Canvas (React Flow) │ Network Map (Hubble UI style) │  │
  │  │  - Dagre hierarchical layout     │  - сокеты, привязанные к      │  │
  │  │  - ego-graph фокус на узле       │    exec_id процессов          │  │
  │  │  - подсветка Attack Path (MITRE) │  - статусы: listen / egress   │  │
  │  ├──────────────────────────────────┴───────────────────────────────┤  │
  │  │ 3-Lane Synchronized Timeline (vis-timeline + timeline-arrows)    │  │
  │  │  Lane 1: Процессы (exec/exit) | Lane 2: Сокеты | Lane 3: Алерты  │  │
  │  │  Sub-second Time-Window Brushing (выделение интервала мышью)     │  │
  │  ├──────────────────────────────────────────────────────────────────┤  │
  │  │ Context Drawer: Incident Summary (Copy SOC) + TracingPolicy YAML │  │
  │  └──────────────────────────────────────────────────────────────────┘  │
  └────────────────────────────────────────────────────────────────────────┘
```

### 3.2. Математическая модель графа (DAG) и сетевой корреляции
1. **Узлы графа процессов ($V_p$):**  
   Каждый узел $v \in V_p$ идентифицируется строкой `process.exec_id` (уникальный хеш Tetragon в base64).
   - Атрибуты: `binary`, `arguments`, `cwd`, `uid`, `pid`.
   - Флаг `in_init_tree`: булево значение, позволяющее отделить штатный контейнерный демон от интерактивной сессии `kubectl exec`.
   - Временные границы: `startTime`, `lastSeen`.
   - Вектор угроз: `mitre: string[]` (`["TA0011", "T1571"]`).
2. **Направленные рёбра графа процессов ($E_p$):**  
   Ребро $(u, v) \in E_p$ создается исключительно по условию:
   $$\text{target.parent\_exec\_id} == \text{source.exec\_id}$$
   Это гарантирует математическую ацикличность графа ($G = (V_p, E_p)$ — строгий DAG) и защищает аналитика от коллизий Linux PID.
3. **Сетевые связи и сокеты ($E_{net}$):**  
   Извлекаются из массива аргументов `process_kprobe.args[0].sock_arg` для ядерных функций:
   - `tcp_connect` / `sys_enter_connect` — исходящие клиентские запросы (`curl`, `wget`).
   - `tcp_close` — завершение сессии.
   - `inet_csk_listen_start` — открытие слушающего сокета на порту (`nc -vv -l -p 8888`).
   - Поля сокета: `saddr:sport` (источник), `daddr:dport` (назначение), `family` (AF_INET), `protocol` (TCP/UDP).
   - Ребро детерминированно привязывается к породившему `exec_id` процесса из блока `process`.
4. **Умное шумоподавление (Smart Noise Folding):**  
   В режиме `All` рутинные вызовы (`ls`, `cat /etc/passwd`, частые healthcheck-пробы) группируются алгоритмом нормализации в компактные узлы-стеки с индикатором «$\times N$ вызовов», которые раскрываются по клику.
5. **Защита от Stored/Reflected XSS:**  
   Злоумышленник внутри контейнера может передать пейлоад в аргументах команды: `nc -l 8888 --comment="<script>alert('XSS')</script>"`.  
   - Frontend рендерит строки исключительно как **Text-as-Data** (стандартное экранирование React без `dangerouslySetInnerHTML`).
   - Backend очищает управляющие символы перед выдачей JSON.

---

## 4. Спецификация интерфейса: Single Pane of Glass (4 квадранта)

Интерфейс строится по эталонному шаблону ведущих систем форензики (Sysdig Inspect, Hubble UI, Falco Sidekick):

### Квадрант 1. Верхняя панель: Фильтры и Instant Highlights (Sysdig Inspect pattern)
- **Селекторы:** Namespace (`hackathon`), Pod (`attacker-pod`), Режим («Все события» / «Только угрозы»).
- **3 KPI-карточки мгновенной сводки:**
  1. *Всего процессов:* общее число узлов в графе.
  2. *Сетевые сокеты:* число активных входящих/исходящих соединений.
  3. *Сработки детекторов (Threat Alerts):* число инцидентов MITRE.
- **Интерактивность:** клик по карточке «Threat Alerts» мгновенно отфильтровывает все квадранты, оставляя на экране только путь атаки.

### Квадрант 2. Левый квадрант: Process Tree Canvas (React Flow + Dagre)
- **Движок:** `@xyflow/react` + `@dagrejs/dagre` с авто-расчетом координат (TB — сверху вниз, 15 строк кода).
- **Кастомные карточки узлов (`.process-node`):**
  - Иконка бинарника (`bash`, `sh`, `curl`, `wget`, `nc`).
  - Метка `kubectl exec` для внешних инъекций против метки `init` для фоновых демонов.
  - Командная строка с кнопкой разворачивания полных аргументов.
  - Цветовая дифференциация: зелёный/серый (норма), оранжевый (подозрительно), пульсирующий красный (инцидент MITRE).
- **Паттерн ego-graph (`checkmarble/ego-graph`):** клик по узлу центрирует холст и скрывает несвязанные ветки, фокусируя внимание на предках и потомках выбранного процесса.

### Квадрант 3. Правый квадрант: Network Topology Map (Hubble UI style)
- **Структура Service Map:** вместо хаотичного облака IP-адресов сокеты группируются в структурированную карту:  
  `Контейнер (attacker-pod)` $\rightarrow$ `Процесс (curl / nc)` $\rightarrow$ `Сервис (target-service:80 / Listen 0.0.0.0:8888)`.
- **Интерактивность:** клик по сокету подсвечивает породивший процесс в левом квадранте и фиксирует момент соединения на таймлайне.

### Квадрант 4. Нижняя сквозная панель: 3-Lane Synchronized Timeline (`vis-timeline`)
- **Многодорожечная структура (Swimlanes):**
  - *Дорожка 1 (Processes):* отрезки реальной длительности жизни (`PROCESS_EXEC` $\rightarrow$ `PROCESS_EXIT`), наглядно показывающие паузы `sleep`.
  - *Дорожка 2 (Network):* точечные маркеры сетевых вызовов (`tcp_connect`, `inet_csk_listen_start`).
  - *Дорожка 3 (Threat Alerts):* маркеры сработок детекторов угроз.
- **Связующие стрелки (`timeline-arrows`):** соединительные линии между дорожками во времени, визуализирующие факт порождения дочернего процесса родителем.
- **Time-Window Brushing:** выделение временного окна мышью мгновенно изолирует верхние графы в границах выбранного отрезка времени (Sub-second drill-down).

### Боковая панель: Context & Response Drawer
- Детальный JSON-инспектор атрибутов: `exec_id`, `parent_exec_id`, `cwd`, `uid`, `arguments`.
- **Кнопка 1 («Скопировать отчет для руководства»):** формирует структурированное резюме инцидента для CISO.
- **Кнопка 2 («Сгенерировать политику TracingPolicy»):** в один клик формирует готовый YAML-манифест Tetragon для блокировки выявленной угрозы (паттерн Tetragon Policy Builder):
  ```yaml
  apiVersion: cilium.io/v1alpha1
  kind: TracingPolicy
  metadata:
    name: block-unauthorized-listen
  spec:
    kprobes:
      - call: "inet_csk_listen_start"
        syscall: false
        args:
          - index: 0
            type: "sock"
        selectors:
          - matchArgs:
              - index: 0
                operator: "Equal"
                values: ["8888"]
            matchActions:
              - action: Sigkill
  ```

---

## 5. Распределение ролей и зон ответственности (3 инженера)

Синхронизация каждые **45–60 минут**. Хранитель API-контракта и интегратор — **Backend Engineer**.

### 5.1. Frontend Engineer
- **Зона:** `apps/web/**`
- **Стек:** React 18+, TypeScript, Vite, `@xyflow/react`, `@dagrejs/dagre`, `vis-timeline`, `zustand`, `lucide-react`.
- **Задачи:**
  - [ ] **F1:** Scaffold Vite React TS, подключение дизайн-системы `activity-map.css`.
  - [ ] **F2:** Реализация Single Pane of Glass (4 квадранта) с заголовком Instant Highlights.
  - [ ] **F3:** Process Tree на `@xyflow/react` + Dagre с кастомными карточками и логикой ego-graph.
  - [ ] **F4:** Network Map (Hubble UI style) с привязкой сокетов к `exec_id`.
  - [ ] **F5:** 3-Lane Timeline на `vis-timeline` (groups). Стрелки `timeline-arrows` — Could.
  - [ ] **F6:** Zustand: Brushing & Linking + Time-Window Brush (Must/Should).
  - [ ] **F7:** Should: Attack Path + MITRE; Could: Replay scrubber / ego-graph package.
  - [ ] **F8:** Should: Incident Summary copy; Could: TracingPolicy YAML; Should: Import JSON dump (выгрузки Events).

### 5.2. Backend Engineer (Интегратор)
- **Зона:** `apps/bff/**`, корень репозитория, целостность API-контрактов.
- **Стек:** Go 1.22+, `net/http` / `chi`, `encoding/json`.
- **Задачи:**
  - [ ] **B1:** Scaffold Go BFF (`apps/bff`), `/healthz`, `/api/activity` с заголовком `X-Data-Source: live | fixture`.
  - [ ] **B2:** Контракт данных: строгий DAG по `exec_id` / `parent_exec_id`, сокеты из `sock_arg`, флаги `in_init_tree`.
  - [ ] **B3:** Парсер-нормализатор Tetragon (`fixtures/raw/events.json` / live): рёбра только `parent_exec_id`→`exec_id`; merge EXEC+EXIT → range; sock_arg → net; **фильтр по `time`/`start_time`** (отрезать bootstrap); поля `Severity`/`Reason`/`TacticsCovered`.
  - [ ] **B4:** Smart Noise Folding (Should) для режима `All`.
  - [ ] **B5:** Live: `public-api`/`history-api` (+ опц. webhook) с fallback fixture; артефакт curl из P6 обязателен до live.
  - [ ] **B6:** XSS sanitize + `go test ./...`.

### 5.3. Platform Engineer
- **Зона:** `scripts/**`, `deploy/**`, `fixtures/raw/**`, кластер k8s и Runtime Radar.
- **Инструменты:** Linux, kind/k3s, Helm 3, eBPF/BTF, `kubectl`.
- **Задачи:**
  - [ ] **P1:** Linux-стенд (не голый Mac Docker Desktop); BTF check; Helm **v0.3.0** + `global.administrator.*`; без prometheus/grafana; port-forward `:9000` / NodePort `:32000`.
  - [ ] **P2:** Sources: process + **incoming socket listen**; Allow `hackathon`; **All**; Apply; smoke Events.
  - [ ] **P3:** `scripts/gen-activity.sh` (4 фазы + **отдельный ls после nc**).
  - [ ] **P4:** Dump `fixtures/raw/events.json` в первый час.
  - [ ] **P5:** 10-мин triage RR → fixture; не чинить eBPF всю ночь.
  - [ ] **P6:** Таблица URL/auth из DevTools + curl для Backend до T4.

---

## 5.4. MoSCoW на 8 часов (чтобы 55% не утонули в киллер-фичах)

| Приоритет | Что | Зачем |
|-----------|-----|--------|
| **Must** | 3 панели (tree/net/timeline) + Zustand sync + fixture/dump + Go `/api/activity` + exec_id DAG | 30%+25% |
| **Should** | Attack Path + MITRE badges + Incident Summary copy + Noise Folding | 20% инновация |
| **Could** | TracingPolicy YAML generator, Replay scrubber, ego-graph npm, timeline-arrows, live webhook | 10% потенциал / вау |
| **Won't** | Писать свои Wasm-детекторы TinyGo, чинить eBPF всю ночь, Cytoscape «на миллионы рёбер» | время |

---

## 5.5. Аудит материалов (docx) + GitHub — что ещё добить

Проверено по `Материалы для карты…docx`, `Материалы по Runtime Radar.docx`, [detectors/guide.md](https://github.com/Runtime-Radar/runtime-radar/blob/main/docs/guides/detectors/guide.md), [Tetragon Events](https://tetragon.io/docs/concepts/events/).

| Тема | Статус в плане | Добивка |
|------|----------------|---------|
| exec_id DAG, Go BFF, React Flow+Dagre, vis-timeline, Zustand, nc:8888, MITRE, BASE.md | Учтено | — |
| **Helm chart version / admin keys** | Слабо | T0: `oci://…/runtime-radar:v0.3.0` (или актуальный tag) + `global.administrator.*`; Habr `v0.1.0` + `auth-center.administrator.*` — **устарело** с v0.2.0. Не деплоить prometheus/grafana без нужды (RAM). |
| **macOS / Docker Desktop** | Слабо | Tetragon на Mac через обычный Docker Desktop **не** взлетит. Нужен **один Linux-стенд** (Lima/`limactl` или чужой Linux) на всю команду; kind — только если `/sys/kernel/btf/vmlinux` доступен в node. |
| **Sources: incoming listen socket** | Слабо | В UI явно включить source **Opening of a socket for incoming connections** (+ process), иначе `inet_csk_listen_start` / детектор nc не придёт. |
| **Flush после nc** | Частично | По detector guide: после `nc -vv -l -p 8888` отдельный `ls -al`, чтобы сбросить буфер в БД. |
| **start_time / bootstrap Tetragon** | Слабо | При старте агент шлёт exec на уже живые процессы. Нормализатор BFF: резать события **до** окна тестовой активности по `time`/`start_time`. |
| **Выгрузки Events (ТЗ)** | Частично (fixture) | Must: dump → `fixtures/raw/events.json`. Should: кнопка **Import JSON** на фронте рядом с live. |
| **Live endpoints** | Задача P6 | Артефакт: таблица URL/method/auth из DevTools `radar-ui` → Events + один рабочий `curl` в README до T4. |
| **Поля угрозы** | Частично | В контракте ждать `Severity`, `Reason`, `TacticsCovered` (не только «mitre[]» на узле). |
| **timeline-arrows / ego-graph** | Как must | Перевести в **Could**; fallback — groups без стрелок / ручной focus без npm ego-graph. |
| **Публичность 6 мес / tie-break** | Нет | DoD: Public + не удалять ≥6 мес; при ничьей выше балл по критерию «Решение» 30%. |

---

## 6. Спецификация скрипта 4-фазной активности Platform

Скрипт `scripts/gen-activity.sh` покрывает все требования задачи PT и создает идеальный визуальный контраст на карте:

```bash
#!/usr/bin/env bash
set -e

echo "[*] 0. Preparing namespace and pods..."
kubectl create namespace hackathon --dry-run=client -o yaml | kubectl apply -f -
kubectl run target-service -n hackathon --image=nginx:alpine --port=80 --expose --dry-run=client -o yaml | kubectl apply -f -
kubectl run attacker-pod -n hackathon --image=alpine:latest --restart=Always -- /bin/sh -c "sleep 3600" || true

echo "[*] Installing forensic & network tools..."
kubectl exec -n hackathon attacker-pod -- apk add --no-cache curl wget netcat-openbsd

echo "[*] ФАЗА 1: Инициализация и разведка (Reconnaissance)..."
# Создает веерное ветвление (sh -> ls, cat, id) и наглядные ступени sleep на таймлайне
kubectl exec -n hackathon attacker-pod -- /bin/sh -c "ls -la /; sleep 2; cat /etc/passwd; sleep 1; id"

echo "[*] ФАЗА 2: Глубокое вложенное ветвление (Deep Process Tree)..."
# Формирует 4-уровневое дерево родственных связей (sh -> sh -> sh -> whoami) для проверки укладки Dagre
kubectl exec -n hackathon attacker-pod -- /bin/sh -c "sh -c 'sleep 1 && ls /tmp && sh -c \"whoami && sleep 2\"'"

echo "[*] ФАЗА 3: Серия сетевых соединений (Network Graph)..."
# Генерирует события tcp_connect / tcp_close из PROCESS_KPROBE для curl и wget
kubectl exec -n hackathon attacker-pod -- /bin/sh -c "
  curl -s http://target-service.hackathon.svc.cluster.local > /dev/null; \
  sleep 2; \
  wget -qO- http://target-service.hackathon.svc.cluster.local > /dev/null
"

echo "[*] ФАЗА 4: Имитация инцидента ИБ (Threat Trigger: MITRE TA0011 / T1571)..."
# inet_csk_listen_start + binary */nc + sport 8888 → CRITICAL (см. detectors/guide.md)
kubectl exec -n hackathon attacker-pod -- /bin/sh -c "nc -vv -l -p 8888 &"
sleep 2
# Отдельный ls — flush буфера событий в БД (обязательно по guide PT)
kubectl exec -n hackathon attacker-pod -- ls -al

echo "[+] Тестовая активность успешно сгенерирована! Проверьте Runtime → Events (threat filter)."
```

**Preflight Platform перед скриптом:**
1. `test -e /sys/kernel/btf/vmlinux` (на Linux-стенде / внутри kind node).  
2. Helm: chart ≥ v0.2.0 keys `global.administrator.*`; предпочтительно **v0.3.0**.  
3. UI Sources: process + **Opening of a socket for incoming connections**; Allow `hackathon`; **All**; Apply.  
4. После скрипта: Events видны; при необходимости Filters → threat events.  


---

## 7. Почасовой план ночи разработки (00:00 — 08:00)

```
00:00 - 01:00 (Т0):  Старт. Инфраструктура кластера, каркасы Web и BFF, первая фикстура.
01:00 - 02:00 (Т1):  Выгрузка дампа событий, Dagre Tree Layout, парсер exec_id -> DAG.
02:00 - 04:00 (Т2):  Network Topology, 3-Lane vis-timeline, Brushing & Linking (Zustand).
04:00 - 05:00 (Т3):  Киллер-фичи: Attack Path Highlight, Replay Scrubber, TracingPolicy Generator.
05:00 - 06:00 (Т4):  Live RR интеграция, Fallback тестирование, Anti-XSS аудит, E2E smoke.
06:00 - 07:00 (Т5):  Заполнение BASE.md (ИИ-декларация), README, CHANGELOG, проверка секретов.
07:00 - 08:00 (Т6):  Feature Freeze, публикация проекта на Мос.Хаб, 2 прогона 3-минутного питча.
```

### Подробные задачи по фазам:

- **Т0 (00:00 — 01:00):**  
  - Platform поднимает kind-кластер с `/sys/kernel/btf/vmlinux`, разворачивает Helm Runtime Radar, проверяет логирование.
  - Backend запускает Go BFF (`:8000`), реализует `/healthz`, `/api/activity` и отдает `sample-activity.json`.
  - Frontend собирает Vite + React + TS, подключает `activity-map.css`, верстает 4-квадрантную сетку.
- **Т1 (01:00 — 02:00):**  
  - Platform настраивает Allow-фильтр в режиме **All**, запускает `gen-activity.sh`, снимает дамп `fixtures/raw/events.json`.
  - Backend пишет нормализатор DAG по `exec_id` и `parent_exec_id`.
  - Frontend реализует компонент Process Tree (`@xyflow/react` + Dagre), выводит карточки процессов с аргументами.
- **Т2 (02:00 — 04:00):**  
  - Frontend интегрирует `vis-timeline` (3 дорожки, `timeline-arrows`) и сетевой граф Hubble-стиля.
  - Frontend настраивает стор Zustand: двусторонняя синхронизация дерева, сети и таймлайна (Brushing & Linking).
  - Backend извлекает параметры сокетов из `sock_arg` и связывает их с `exec_id`.
  - Platform снимает маршруты к `public-api` через DevTools Network.
- **Т3 (04:00 — 05:00):**  
  - Реализация кнопки «Подсветить цепочку атаки» (Attack Path) с анимацией красного градиента.
  - Добавление Time-Travel ползунка для перемотки состояния контейнера во времени.
  - Реализация бокового дровера: генератор Incident Summary для SOC + генератор YAML `TracingPolicy`.
- **Т4 (05:00 — 06:00):**  
  - Подключение live API Runtime Radar с автоматическим fallback на fixture (`X-Data-Source: live | fixture`).
  - Проверка безопасности: санитизация строк от XSS, проверка отсутствия утечек секретов.
- **Т5 (06:00 — 07:00):**  
  - Совместно: написание обязательного `BASE.md` (архитектура, описание использования ИИ, подтверждение проверки безопасности кода).
  - Оформление `README.md` (инструкция запуска за 2 команды) и `CHANGELOG.md`.
- **Т6 (07:00 — 08:00):**  
  - 07:00 — Полный Code Freeze.
  - 07:15 — Публикация репозитория на платформе «Мос.Хаб».
  - 07:30 — Два полноценных прогона 3-минутного выступления перед жюри с секундомером.

---

## 8. Сценарий защиты перед жюри: «15 секунд расследования» (3 минуты)

```
00:00 - 00:30  Проблема: Ночной звонок руководства и хаос плоских таблиц логов
00:30 - 01:45  Живое Демо: 15-секундное расследование инцидента на 4 квадрантах
01:45 - 02:30  Техническое превосходство: eBPF, Tetragon, строгий DAG exec_id, Anti-XSS
02:30 - 03:00  Потенциал внедрения: интеграция в radar-ui и 1-click TracingPolicy
```

### Текст выступления спикера:

**00:00 — 00:30 | Боль и Проблема**
> *«Здравствуйте, уважаемые эксперты! Представьте реальность аналитика SOC: 3 часа ночи, срабатывает алерт. Руководитель подходит с вопросом: "Что произошло в контейнере?". В руках у безопасника — плоская таблица из 20 тысяч строк в режиме All. PID переиспользованы ядром, дочерние вызовы перемешаны, сетевые соединения оторваны от процессов. Анализ занимает часы.*  
> *Команда OnlyGo разработала Интерактивную карту активности контейнера — инструмент, дающий полный ответ за 15 секунд в одном окне».*

**00:30 — 01:45 | Живое Демо (Экран Single Pane of Glass)**
> *«Наш интерфейс построен по концепции Single Pane of Glass:*  
> *1. **(Взгляд на Instant Highlights):** Сверху — мгновенная сводка. Кликаем на плитку "Threat Alerts" — и экран мгновенно очищается от рутинного шума.*  
> *2. **(Клик по Attack Path):** Нажимаем кнопку "Цепочка атаки". Система красным градиентом подсвечивает путь вторжения от внешней инъекции sh -c через разведку к слушающему бэкдору nc на порту 8888.*  
> *3. **(Зум в Process Tree):** Никаких коллизий PID! Наше дерево строится по кластерным exec_id от Tetragon с авто-укладкой по алгоритму Dagre.*  
> *4. **(Клик по процессу):** Срабатывает Brushing & Linking: на сетевой карте Hubble-стиля вспыхивает сокет 8888, а на нижнем таймлайне виден точный момент сработки детектора inet_csk_listen_start с бейджами MITRE ATT&CK TA0011 и T1571.*  
> *5. **(Клик Time-Window Brushing):** Выделяем мышью двухсекундный интервал на таймлайне — и верхние графы показывают состояние контейнера с субсекундной точностью.*  
> *6. **(Клик в Response Drawer):** Один клик — и структурированный отчет для CISO готов в буфере обмена. Второй клик — сгенерирован YAML манифест TracingPolicy для мгновенной блокировки угрозы на уровне ядра eBPF. Ровно 15 секунд!»*

**01:45 — 02:30 | Инженерная реализация и безопасность**
> *«Под капотом — нативный стек Runtime Radar. Сенсор Tetragon собирает события. Наш Go BFF нормализует поток в строгий ациклический граф по exec_id и извлекает сетевые структуры sock_arg.*  
> *Особое внимание мы уделили требованиям безопасной разработки: строки командных строк контейнера экранируются от Stored XSS-атак на аналитика SOC.*  
> *Архитектура двухуровневая: при потере связи с кластером система бесшовно переключается на офлайн-снапшот с индикатором X-Data-Source».*

**02:30 — 03:00 | Потенциал и финал**
> *«Наше решение спроектировано как модульное расширение для штатного radar-ui и готово к работе с public-api и ClickHouse history-api.*  
> *OnlyGo превращает терабайты шума в мгновенное понимание инцидента. Спасибо!»*

---

## 9. Чек-лист сдачи проекта на «Мос.Хаб» (Definition of Done)

### Функционал и качество (55% оценки)
- [ ] Все 4 квадранта (Header KPI, Process Tree, Network Map, 3-Lane Timeline) работают в **едином окне**.
- [ ] Дерево процессов строится строго как DAG по `exec_id` и `parent_exec_id` (нет коллизий PID).
- [ ] Сетевые сокеты привязаны к породившим их `exec_id` процессов.
- [ ] Реактивный Brushing & Linking синхронизирует выбор между деревом, сетью и таймлайном.
- [ ] Работает Sub-second Time-Window Brushing на шкале времени.
- [ ] Реализована подсветка Attack Path с бейджами MITRE ATT&CK (TA0011 / T1571).
- [ ] Go BFF корректно отдает `X-Data-Source: live | fixture` и защищен от падений.

### Инновации и потенциал (30% оценки)
- [ ] 1-Click SOC Incident Summary (копирование резюме для руководства).
- [ ] 1-Click TracingPolicy Generator (генерация YAML политики блокировки для Tetragon).
- [ ] Time-Travel Scrubber (Replay) для пошаговой перемотки инцидента.
- [ ] Архитектурная совместимость с экосистемой Runtime Radar (`radar-ui`, `public-api`, ClickHouse).

### Регламент и безопасность (15% оценки)
- [ ] Репозиторий опубликован как **публичный** проект на «Мос.Хаб» (не удалять ≥ **6 месяцев**).
- [ ] В корне **`BASE.md`**: архитектура, ИИ по этапам, XSS/секреты проверены (шаблон `hackathon-kit/stubs/BASE.md`).
- [ ] Проведен аудит безопасности: строгая защита от Stored XSS в выводе команд, отсутствие секретов/паролей в истории коммитов Git (`08-git-secrets.md`).
- [ ] Подготовлена презентация (PPTX/PDF) по шаблону финалиста.
- [ ] Отрепетирован 3-минутный питч со сценарием «15 секунд расследования».
