# Полный план команды OnlyGo — «Ночь кода 2026»

**Задача:** Интерактивная карта активности контейнера (Positive Technologies / Runtime Radar)  
**Команда:** **3 человека** (у каждого свой список задач)  
**Репо сдачи:** `OnlyGo` на Мос.Хаб (`git@hub.mos.ru:event_DIT/2026/onlygo.git`)  
**Официальное описание:** `hackathon-kit/refs/Задача_Positive_Technologies_описание_задачи.pdf`

Эксперт на площадке (Святослав Сорокин, PT / Runtime Radar): цель — помочь аналитику ИБ **быстро восстановить картину** в контейнере: порядок действий, дерево процессов, события, сеть — **в одном UX**, а не в «диком списке» Events.

---

## 0. Жёсткое правило организаторов

Код решения — **только после старта**. Жюри смотрит историю коммитов.  
Локально держим `hotel/` + `hackathon-kit/` + preview; в OnlyGo до старта — не пушить apps/AGENTS/skills.

---

## 1. Что сказал эксперт → как закрываем

### Продукт (слайд «О задаче»)

Веб-интерфейс по данным Runtime Radar: в **едином** окне  
хронология + дерево процессов (ppid→pid) + события + сеть.  
Пользователь: аналитик ИБ. Метрика питча: **time-to-insight < 60 с**.

### Необходимые компетенции (слайд «НЕОБХОДИМЫЕ») → роли

| # | Компетенция эксперта | Кто у нас |
|---|----------------------|-----------|
| 01 Фронтенд | веб / UI | **A — Frontend** |
| 02 Визуализация | графы, деревья, таймлайны; **готовые либы ok** | **A — Frontend** (React Flow + dagre) |
| 03 Данные | работа с выгрузками событий | **B — Data / BFF** |
| 04 Kubernetes | базовая контейнеризация / kind / Helm | **C — Platform + Docs** |

### Подготовка данных RR (слайд «ПОДГОТОВКА ДАННЫХ») — чеклист C (+ B принимает dump)

1. **Развернуть Runtime Radar** — Helm в kind/minikube/k3s, eBPF/BTF; docs quickstart.  
2. **Настроить сбор** — Sources; Runtime → Parameters → Allow filter на ns (напр. `hackathon`); сохранение **All**; Apply.  
3. **Тестовая активность** — поды в `hackathon`; `kubectl exec`: `ls`/`cat`/`sleep`; дети через `sh -c`; повторы `curl`/`wget` с паузами.  
4. **Проверить** — Runtime → Events: тестовые контейнеры и детали видны.  

У нас локально RR уже поднимали; ночью C повторяет 2–4 и кормит B свежим dump. Fallback: `hackathon-kit/fixtures/sample-activity.json`.

---

## 2. Три человека — списки задач

Sync каждые **45–60 мин**, 5 минут. Integrator контракта = **B**. Корень репо = **B** (или C по просьбе B).

### A — Frontend + визуализация

**Зона:** только `apps/web/**`  
**Skill:** `hotel/skills/03-frontend-activity-map.md` (+ `09` для Incident Summary)

| # | Задача | AC |
|---|--------|-----|
| A1 | Vite React TS scaffold + `activity-map.css` из kit | UI открывается |
| A2 | Типы из OpenAPI BFF / `stubs/types.ts` | fetch `/api/activity` |
| A3 | Process tree на **@xyflow/react + dagre** | ppid→pid, роли цветов |
| A4 | Network panel | edges видны, клик выделяет |
| A5 | Timeline | события по времени |
| A6 | Selection sync между тремя панелями | один selected id |
| A7 | **Incident Summary** + Copy | текст SOC по шаблону |
| A8 | Replay (если успеем) | scrub по таймлайну |
| A9 | Не трогать BFF/корень; pacing коммитов | см. §4 |

**Не делает:** Helm/RR UI-настройку, Pydantic, слайды (может помочь C 10 мин на скрины).

---

### B — Данные / BFF

**Зона:** `apps/bff/**`, fixture в bff, OpenAPI  
**Skill:** `hotel/skills/02-backend-activity-bff.md`

| # | Задача | AC |
|---|--------|-----|
| B1 | FastAPI + Pydantic `ActivityPayload` | `/openapi.json` |
| B2 | `GET /api/activity` + чтение fixture | `X-Data-Source: fixture` |
| B3 | Normalize dump от C → наша схема | pytest схемы зелёный |
| B4 | Live-клиент RR HTTP `:9000` (бонус) | fallback на fixture при ошибке |
| B5 | Integrator: мерж MR A/C по контракту | smoke curl + UI |
| B6 | Корень: docker-compose/workspaces только B | без секретов |

**Не делает:** React Flow, долгий тюнинг kind (это C).

---

### C — Platform (K8s/RR) + питч + changelog/docs

**Зона:** `scripts/**`, `deploy/**`, `fixtures/raw/**`, `docs/**`, `CHANGELOG.md`, `README.md`, pptx  
**Skills:** `01-platform-rr.md`, `04-pitch-demo.md`, `08-git-secrets.md`, `09-soc-demo-language.md`, `10-docs-changelog.md`

| # | Задача | AC |
|---|--------|-----|
| C1 | RR up: port-forward `:9000`, pods Ready | UI логин |
| C2 | Sources + Allow filter ns `hackathon` (или default) + All | Events не пустые |
| C3 | Скрипт тестовой активности (exec + curl loop) | дерево/сеть в Events |
| C4 | Dump сырых событий → `fixtures/raw/` для B | файл + короткий note |
| C5 | Fixture fallback всегда под рукой | демо без кластера |
| C6 | **docs/** + **CHANGELOG.md** (правило §5) | после каждой крупной фичи |
| C7 | README: 3 команды запуска | жюри поднимает |
| C8 | **`BASE.md`** — базовый файл сдачи (+ блок ИИ) | см. §2.1 |
| C9 | Презентация + демо-скрипт < 60 с | репетиция 2× |
| C10 | Freeze-чеклист с A/B (`05-freeze-review`) | Public visibility |

**Не делает:** компоненты xyflow, модели Pydantic (кроме передачи dump).

### 2.1 Формат сдачи (из PDF PT)

Публичный проект на Мос.Хаб должен содержать:

| Артефакт | Кто | Где в OnlyGo |
|----------|-----|--------------|
| Исходный код веб-интерфейса (+ BFF) | A, B | `apps/web`, `apps/bff` |
| Документация | C | `docs/`, `README.md` |
| **Базовый файл** | C (факты от A/B) | **`BASE.md`** в корне |
| Презентация | C | pptx / `docs/pitch/` |
| Демо жюри | все, ведёт C | сценарий `docs/DEMO.md` |

**`BASE.md` (обязательно по PDF)** — текстовый файл о ПО:

- что сделано и для кого (аналитик ИБ / Runtime Radar);
- стек и как запустить;
- **какие ИИ-инструменты использовались и на каких этапах**;
- что ключевые решения осмыслены командой;
- для ИБ: сгенерированный код просмотрен (секреты, небезопасные практики).

Шаблон: `hackathon-kit/stubs/BASE.md` → копировать в OnlyGo после старта.  
Правила ведения: `hotel/skills/10-docs-changelog.md`.

Общие рамки PDF: команда 3–5 чел., ≤8 часов на площадке, ИИ разрешён и приветствуется; права/публичность — по регламенту Мос.Хаб.

---

## 3. Продуктовая рамка и критерии

| Критерий | Вес | Как |
|----------|-----|-----|
| Решение | 30% | одно окно: timeline + tree + network |
| Качество | 25% | UI, контракт, README, fixture |
| Инновация | 20% | Incident Summary (+ Replay) |
| Презентация | 15% | C ведёт; A/B на демо у экрана |
| Потенциал | 10% | Now/Next/Later в слайдах |

**Питч:** не «красивый граф», а «сокращаем время расследования для SOC».

---

## 4. Агенты, зоны, pacing

Правила — только в `МосХаб/hotel/` (не в OnlyGo).

Промпт агенту:  
`Разрешённые пути: … . Чужое read-only. Не меняй корень без роли B.`

Стек: Web = Vite/React/xyflow/dagre · BFF = FastAPI/Pydantic · fixture-first.

Коммиты дробить: scaffold → fixture → tree → timeline/network → sync → summary → live → docs.

---

## 5. Документация и CHANGELOG (как идея gotender, упрощённо)

В OnlyGo **после старта** (пишет **C**, факты дают A/B):

```
docs/
  README.md          # оглавление docs
  DEMO.md            # сценарий жюри < 60 с
  architecture.md    # 10–20 строк: RR → BFF → Web
CHANGELOG.md         # журнал изменений ночи (корень)
BASE.md              # обязательный «базовый файл» сдачи (ИИ + описание)
reports/             # опционально короткие SUMMARY_*.md после крупного куска
```

**Правило changelog (обязательно):**

1. После мержа осмысленного этапа (не каждый микрокоммит) — запись в `CHANGELOG.md`:  
   дата/время · кто (A/B/C) · что сделано · как проверить.  
2. Conventional commits в git (`feat(web):`, `feat(bff):`, `docs:`, `fix:`).  
3. Крупный кусок (tree готов / live RR / summary) → опционально `reports/SUMMARY_<тема>_YYYY-MM-DD.md` на 5–15 строк.  
4. Перед freeze — заполнить **`BASE.md`** (ИИ-этапы актуальны).  
5. Не писать секреты RR в docs / BASE.  
6. AGENTS/skills в docs **не** класть.

Шаблоны: `hackathon-kit/stubs/docs/`, `stubs/CHANGELOG.md`, `stubs/BASE.md` — копировать в OnlyGo после старта.

---

## 6. Архитектура

```
Runtime Radar (:9000) ──▶ BFF (Pydantic/OpenAPI) ──▶ Web (xyflow + panels + Incident Summary)
         │                        ▲
         └── dump / fixture ──────┘
```

Live: HTTP API, не парсинг логов. См. `.rr-install/ACCESS.md`.

---

## 7. Чеклист до старта

- [ ] Docker / kind / RR / Events (C)  
- [ ] Fixture + design pack на диске  
- [ ] Роли A/B/C понятны всем  
- [ ] Ноль кода решения в OnlyGo до сигнала  
- [ ] hotel/ открыт у агентов локально  

---

## 8. План ночи

### T0 (20–30 мин)
- Pull OnlyGo; hotel локально  
- B: scaffold BFF + Pydantic  
- A: scaffold web + css  
- C: RR filter + gen-activity; скопировать `stubs/docs` → `docs/`, `stubs/CHANGELOG.md`, `stubs/BASE.md`

### T1–T2
- B: fixture API + pytest  
- A: tree xyflow  
- C: dump для B; первая запись changelog

### T2–T4
- A: network + timeline + sync  
- B: normalize (+ live если успеет)  
- C: DEMO.md + слайды каркас

### T4–T6
- A: Incident Summary (+ Replay)  
- B: стабильный fallback  
- C: README, changelog, черновик BASE.md (ИИ), репетиция

### T6–T7 Freeze
- E2E на одном ноуте; дописать BASE.md; Public; питч 2×

---

## 9. Риски

| Риск | Митигация |
|------|-----------|
| 3 человека, мало рук на питч | C совмещает docs+питч; A/B не отвлекаются |
| RR/eBPF на kind | fixture-first |
| Рассинхрон FE/BE | OpenAPI с T0, integrator B |
| Нет истории для жюри | changelog + дробные коммиты |

---

## 10. DoD

- [ ] Tree + network + timeline + sync  
- [ ] Incident Summary  
- [ ] Fixture и/или live  
- [ ] README + CHANGELOG + **BASE.md** + презентация  
- [ ] Демо < 60 с  
- [ ] Public, без секретов и без hotel в git  

---

## 11. Ссылки

- Hotel: `../hotel/`  
- Design: `design/activity-map.css`  
- Fixture: `fixtures/sample-activity.json`  
- Docs stubs: `stubs/docs/`, `stubs/BASE.md`, `stubs/CHANGELOG.md`  
- PDF задачи: `refs/Задача_Positive_Technologies_описание_задачи.pdf`  
- RR: `../.rr-install/ACCESS.md`  
