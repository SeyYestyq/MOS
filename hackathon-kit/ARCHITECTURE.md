# Архитектура продукта — вектор OnlyGo

**Что строим:** веб-модуль «карта активности контейнера» — один экран расследования инцидента по данным Runtime Radar.  
**Не строим:** свой eBPF/Tetragon, клон radar-ui, ML-платформу, Wasm-детекторы.

Цель продукта: аналитик за **~15 секунд** отвечает «что / как / почему» без листания тысяч строк Events.

Код сдачи — только репозиторий **OnlyGo**. Скиллы — в `hotel/` рядом на диске.

---

## 1. Вектор в одном абзаце

Плоский поток Tetragon (REST **или** сохранённый dump) проходит **один адаптер** и становится одним снимком. Из него UI берёт три проекции: процессы, сетевые наблюдения, события и угрозы. Живой RR желателен; **fixture-first** обязателен. Канон полей: `DATA_MODEL.md`.

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
│  ОДИН адаптер: API и dump → одни события                    │
│  снимок: процессы + net observations + threats              │
│  GET /api/activity   X-Data-Source: live | raw | fixture    │
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

| Кто | Зона | Стремится к |
|-----|------|-------------|
| **Ангелина** | `apps/web` | одно окно, sync панелей, Attack Path |
| **Денис** | `apps/bff` + контракт + тесты | DAG exec_id + sock_arg + fixture + `go test` |
| **Дамир** | RR / scripts / raw dump + live-стык | живые события, dump с args, сценарий nc:8888 |

Чеклисты: `TEAM_ROLES.md`.

---

## 3. Модель данных (ядро продукта)

Канон: **`DATA_MODEL.md`** (модель напарника, принята). Кратко:

| Проекция | Правило |
|----------|---------|
| **Процессы** | ребро только если родитель **в снимке**; иначе `parentStatus: outside_window`, не клеить к init |
| **Времена** | `startedAt`, подтверждённый `exitedAt`, `firstObservedAt`, `lastObservedAt` (last ≠ exit) |
| **Контейнер** | cluster + namespace + pod name + **pod UID** + container id |
| **Сеть** | наблюдения `connect` / `listen` / `close` из `sock_arg`, не список «живых сокетов» |
| **Угрозы** | отдельные сущности: detector, severity, reason, tactics, techniques, ссылка на событие |
| **Контракт** | один JSON; Go и TS сверяются тестом, не правятся врозь |

Демо-якорь: `nc -l 8888` → listen + threat **TA0011 / T1571**.

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
live REST slice ──┐
                   ├─ один адаптер ─► один JSON (X-Data-Source честный)
сохранённый dump ──┘
fixture — тот же shape, если RR нет
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
