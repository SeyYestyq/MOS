# Архитектура продукта — вектор OnlyGo

**Что строим:** веб-модуль «карта активности контейнера» — один экран расследования инцидента по данным Runtime Radar.  
**Не строим:** свой eBPF/Tetragon, клон radar-ui, ML-платформу, Wasm-детекторы.

Цель продукта: аналитик за **~15 секунд** отвечает «что / как / почему» без листания тысяч строк Events.

Код сдачи — только репозиторий **OnlyGo**. Скиллы — в `hotel/` рядом на диске.

---

## 1. Вектор в одном абзаце

Плоский поток событий Tetragon (через Runtime Radar) → **Go BFF нормализует** в три согласованных среза данных (процессы / сеть / время) → **React Single Pane of Glass** показывает их синхронно. Живой RR желателен; **fixture-first** обязателен (демо не зависит от кластера).

---

## 2. Слои (что чьё)

```
┌─────────────────────────────────────────────────────────────┐
│  Runtime Radar (чужой продукт, мы потребители)              │
│  Tetragon → event-processor → public-api / dump / webhook   │
│  ⚠ сеть (IP/порты) только из полного JSON (REST), не из CH  │
└────────────────────────────┬────────────────────────────────┘
                             │ JSON events (live | dump | fixture)
                             ▼
┌─────────────────────────────────────────────────────────────┐
│  apps/bff  — Go  :8000          Backend                     │
│  normalize → DAG exec_id · sock_arg → net · XSS sanitize    │
│  GET /api/activity   X-Data-Source: live | fixture          │
└────────────────────────────┬────────────────────────────────┘
                             │ один JSON-контракт (types.ts)
                             ▼
┌─────────────────────────────────────────────────────────────┐
│  apps/web  — React+Vite :5173   Frontend                    │
│  Tree (xyflow+dagre) · Network · Timeline (vis-timeline)    │
│  Zustand: один selectedExecId + time brush на все панели    │
│  Drawer: Incident Summary (+ Could: TracingPolicy YAML)     │
└─────────────────────────────────────────────────────────────┘

Platform: kind/Helm RR · Sources/All · gen-activity · dump → fixtures/raw/
```

| Роль | Зона | Стремится к |
|------|------|-------------|
| **Frontend** | `apps/web` | одно окно, sync панелей, читаемый Attack Path |
| **Backend** | `apps/bff` + контракт | правильный DAG + сеть из sock_arg + fallback |
| **Platform** | кластер / scripts / raw dump | живые события + воспроизводимый сценарий nc:8888 |

---

## 3. Модель данных (ядро продукта)

Не «логи в таблице», а **три проекции одного инцидента**:

| Проекция | Источник в событии | Правило |
|----------|--------------------|---------|
| **Дерево процессов** | `process.exec_id`, `parent.exec_id` | ребро только `parent_exec_id → exec_id` (**не** pid/ppid) |
| **Сеть** | `PROCESS_KPROBE` → `args[].sock_arg` | listen / connect / close, привязка к `exec_id` |
| **Таймлайн** | EXEC+EXIT → интервал; kprobe → точка | те же `execId`, общий brush |

Демо-якорь угрозы: `nc -l 8888` → `inet_csk_listen_start` → MITRE **TA0011 / T1571** (CRITICAL).

Контракт: `stubs/CONTRACT.md` + `stubs/types.ts` → ответ `GET /api/activity`.

---

## 4. UX-цель (как выглядит «готово»)

Один viewport, без вкладок «сначала дерево, потом сеть»:

1. **Highlights** — быстрые фильтры (угрозы / процессы / сокеты).
2. **Process Tree** — иерархия Dagre; init vs `kubectl exec`; подсветка пути атаки.
3. **Network Map** — сокеты того же `exec_id`.
4. **Timeline** — 2–3 дорожки; выделение интервала режет шум наверху.
5. **Incident Summary** — текст для CISO/тикета (copy).

Синхронизация = продукт. Три красивых, но **несвязанных** виджета = провал вектора.

---

## 5. Режимы данных (деградация)

```
live REST slice (с args)  →  dump fixtures/raw  →  sample-activity.json
         ▲                        ▲                       ▲
      Platform                 Platform                 Backend stub
```

Заголовок `X-Data-Source` всегда честный. UI обязан работать на fixture.

---

## 6. Что внутри вектора / что снаружи

| Внутри (цель ночи) | Снаружи (не отвлекает) |
|--------------------|-------------------------|
| 3 панели + Zustand sync | свой Tetragon / TracingPolicy как Must |
| exec_id DAG + sock_arg net | обучение GNN / LoRA / HF-датасеты |
| fixture + REST dump | полный клон radar-ui (Angular) |
| MITRE на демо-цепочке | Cytoscape «на миллионы рёбер» |
| BASE.md + демо &lt; 60 с | Python в рантайме BFF |

**Could (если Must жив):** Noise Folding, Import JSON, TracingPolicy YAML, offline `risk` score на узлах.

---

## 7. Фраза для команды и питча

> Мы не рисуем «ещё один дашборд Events». Мы собираем из Runtime Radar **согласованную карту инцидента**: кто кого породил, кто куда сходил по сети, в каком порядке — и даём это в одном окне за секунды.

Подробный план: `PLAN_NOCH_KODA.md`. Старт переноса: `../hotel/prompts/QUICK_START_FOR_TEAM.md`.
