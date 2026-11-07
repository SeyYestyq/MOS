# Полный план команды OnlyGo — «Ночь кода 2026»

**Задача:** Интерактивная карта активности контейнера (Positive Technologies / Runtime Radar)  
**Команда:** **3 человека** — три инженерные роли (Frontend / Backend / Platform)  
**Репо сдачи:** `OnlyGo` на Мос.Хаб (`git@hub.mos.ru:event_DIT/2026/onlygo.git`)  
**Официальное описание:** `hackathon-kit/refs/Задача_Positive_Technologies_описание_задачи.pdf`

Эксперт (Святослав Сорокин, PT): помочь аналитику ИБ **быстро восстановить картину** в контейнере — порядок действий, дерево процессов, события, сеть в **одном UX**.

Docs / `BASE.md` / питч — **не отдельная роль**: каждый пишет changelog по своим коммитам; `BASE.md` и слайды собирают **все вместе** в конце (20–30 мин).

---

## 0. Жёсткое правило организаторов

Код решения — **только после старта**. Жюри смотрит историю коммитов.  
Локально: `hotel/` + `hackathon-kit/`; в OnlyGo до старта — не пушить apps/AGENTS/skills.

---

## 1. Эксперт / PDF → наши роли

| # эксперта | Компетенция | Роль |
|------------|-------------|------|
| 01 Фронтенд | веб / UI | **Frontend** |
| 02 Визуализация | графы, деревья, таймлайны (готовые либы ok) | **Frontend** (`@xyflow/react` + dagre) |
| 03 Данные | выгрузки событий | **Backend** (BFF + normalize) |
| 04 Kubernetes | kind / Helm / RR | **Platform** |

### Подготовка данных RR — чеклист Platform (Backend принимает dump)

1. Runtime Radar в Kubernetes (Helm; kind/minikube/k3s; eBPF/BTF).  
2. Sources; Runtime → Parameters → Allow filter ns (напр. `hackathon`); сохранение **All**; Apply.  
3. Тестовая активность: pods + `kubectl exec` (`ls`/`cat`/`sleep`/`sh -c`) + `curl`/`wget` с паузами.  
4. Runtime → Events: события видны.  

Fallback: `hackathon-kit/fixtures/sample-activity.json`.

---

## 2. Три человека — списки задач

Sync каждые **45–60 мин**. Integrator контракта = **Backend**.

### Frontend

**Зона:** `apps/web/**`  
**Skills:** `03-frontend-activity-map.md`, `09-soc-demo-language.md`

| # | Задача | AC |
|---|--------|-----|
| F1 | Vite React TS + `activity-map.css` | UI открывается |
| F2 | Типы из OpenAPI / `stubs/types.ts` | fetch `/api/activity` |
| F3 | Process tree — **xyflow + dagre** | ppid→pid |
| F4 | Network panel | edges + select |
| F5 | Timeline | события по времени |
| F6 | Selection sync | один selected id |
| F7 | Incident Summary + Copy | текст SOC |
| F8 | Replay (если успеем) | scrub таймлайна |

**Не делает:** Helm/RR, Pydantic BFF.

---

### Backend

**Зона:** `apps/bff/**` (+ integrator корня)  
**Skill:** `02-backend-activity-bff.md`

| # | Задача | AC |
|---|--------|-----|
| B1 | FastAPI + Pydantic `ActivityPayload` | `/openapi.json` |
| B2 | `GET /api/activity` + fixture | `X-Data-Source: fixture` |
| B3 | Normalize dump от Platform | pytest схемы |
| B4 | Live-клиент RR HTTP `:9000` | fallback на fixture |
| B5 | Integrator / smoke | curl + UI с Frontend |
| B6 | Корень: compose/workspaces | без секретов |

**Не делает:** React Flow; долгий тюнинг kind (это Platform).

---

### Platform

**Зона:** `scripts/**`, `deploy/**`, `fixtures/raw/**`  
**Skills:** `01-platform-rr.md`, `07-focused-fix.md` (если RR падает)

| # | Задача | AC |
|---|--------|-----|
| P1 | RR up, port-forward `:9000` | UI логин |
| P2 | Sources + Allow filter + All | Events не пустые |
| P3 | Скрипт тестовой активности | дерево/сеть в Events |
| P4 | Dump → `fixtures/raw/` для Backend | файл готов |
| P5 | Стабильный demo-path на ночь | команда знает fallback |
| P6 | Помощь Backend с live API (Network tab / endpoints) | список URL в чате |

**Не делает:** xyflow-компоненты; не «человек на слайды» — питч общий.

### Общая лёгкая нагрузка (все трое, не основная работа)

| Артефакт | Как |
|----------|-----|
| `CHANGELOG.md` | каждый дописывает блок после своего крупного этапа |
| `README.md` / `docs/` | Backend — запуск BFF; Frontend — запуск web; Platform — RR/fixture notes |
| **`BASE.md`** | вместе за 20 мин до freeze (шаблон `stubs/BASE.md`) |
| Презентация + демо | вместе; сценарий `docs/DEMO.md`; говорит тот, кто увереннее на сцене |
| Skill docs | `10-docs-changelog.md`, `04-pitch-demo.md`, `08-git-secrets.md` — читать всем |

### 2.1 Формат сдачи (PDF PT)

Публичный проект: код + документация + **`BASE.md`** (базовый файл + ИИ) + презентация + демо.  
Шаблоны: `hackathon-kit/stubs/`. ИИ разрешён; в BASE указать инструменты и этапы.

---

## 3. Критерии жюри

| Критерий | Вес | Как |
|----------|-----|-----|
| Решение | 30% | timeline + tree + network в одном окне |
| Качество | 25% | UI, контракт, README, fixture |
| Инновация | 20% | Incident Summary (+ Replay) |
| Презентация | 15% | общая репетиция |
| Потенциал | 10% | Now/Next/Later |

Питч: «сокращаем время расследования для SOC», не «красивый граф».

---

## 4. Агенты и pacing

Правила только в `МосХаб/hotel/`.  
Промпт: `Разрешённые пути: … . Чужое read-only.`

Стек: Vite/React/xyflow · FastAPI/Pydantic · fixture-first.  
Коммиты: scaffold → fixture → tree → panels → sync → summary → live → docs/BASE.

---

## 5. Документация

```
docs/README.md, DEMO.md, architecture.md
CHANGELOG.md
BASE.md
```

После этапа — запись в CHANGELOG (кто / что / как проверить). Без секретов. Без AGENTS в OnlyGo.

---

## 6. Архитектура

```
Runtime Radar (:9000) ──▶ BFF (Pydantic) ──▶ Web (xyflow + panels + Incident Summary)
         │                     ▲
         └── dump / fixture ───┘
```

Live: HTTP API. См. `.rr-install/ACCESS.md`.

---

## 7. Чеклист до старта

- [ ] Docker / kind / RR / Events (Platform)  
- [ ] Fixture + design pack  
- [ ] Роли Frontend / Backend / Platform понятны  
- [ ] Ноль кода решения в OnlyGo до сигнала  

---

## 8. План ночи

### T0
- Frontend: scaffold web + css  
- Backend: scaffold BFF + Pydantic  
- Platform: filter + gen-activity; stubs docs/CHANGELOG/BASE в репо  

### T1–T2
- Backend: fixture API  
- Frontend: tree  
- Platform: dump  

### T2–T4
- Frontend: network + timeline + sync  
- Backend: normalize (+ live)  
- Platform: держит RR / чинит Events  

### T4–T6
- Frontend: Incident Summary  
- Backend: стабильный fallback  
- Все: README куски  

### T6–T7 Freeze
- E2E; вместе BASE.md + слайды; Public; репетиция 2×  

---

## 9. Риски

| Риск | Митигация |
|------|-----------|
| RR/eBPF | fixture-first |
| Рассинхрон FE/BE | OpenAPI, integrator Backend |
| Docs забыли | 20 мин в конце все трое на BASE |

---

## 10. DoD

- [ ] Tree + network + timeline + sync + Incident Summary  
- [ ] Fixture и/или live  
- [ ] README + CHANGELOG + BASE.md + презентация  
- [ ] Демо < 60 с; Public; без hotel/секретов  

---

## 11. Ссылки

- Hotel: `../hotel/`  
- Design: `design/activity-map.css`  
- Fixture: `fixtures/sample-activity.json`  
- Stubs: `stubs/BASE.md`, `stubs/CHANGELOG.md`, `stubs/docs/`  
- PDF: `refs/Задача_Positive_Technologies_описание_задачи.pdf`  
- RR: `../.rr-install/ACCESS.md`  
