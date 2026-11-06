# Полный план команды OnlyGo — «Ночь кода 2026»

**Задача:** Интерактивная карта активности контейнера (Positive Technologies / Runtime Radar)  
**Команда:** 4 человека  
**Репо сдачи:** `OnlyGo` на Мос.Хаб (`git@hub.mos.ru:event_DIT/2026/onlygo.git`)  
**Сейчас в репо:** только PDF требований/критериев + шаблон презентации  

---

## 0. Жёсткое правило организаторов (читать всем)

Из письма Мос.Хаб:

> Код решения пишем **только на площадке**. Установить и проверить ПО нужно заранее, а писать код задачи — **только после старта** хакатона. Жюри смотрит **историю коммитов** в проекте на Мос.Хаб, поэтому **заготовки, принесённые с собой, будут видны**.

### Что можно ДО старта (7.11)

| Можно | Нельзя |
|--------|--------|
| Ставить Docker, kubectl, helm, kind, браузер, IDE | Пушить код решения в OnlyGo |
| Поднять Runtime Radar локально, проверить Events | Коммитить AGENTS.md / skills / UI / BFF в OnlyGo заранее |
| Читать ТЗ, критерии, презентацию | Копировать готовый `activity-map-preview` в OnlyGo до старта |
| Держать **локально** шпаргалки: `hackathon-kit/` + `hotel/` | Считать kit «уже сданным»; пушить AGENTS/skills в OnlyGo |

**Локальные папки (НЕ в OnlyGo git):**

- `МосХаб/hotel/` — **AGENTS.md, skills, промпты агентов** (агенты читают с диска; в Mos.Хаб не коммитим)  
- `МосХаб/hackathon-kit/` — дизайн, fixture, план, черновик контракта  
- `МосХаб/activity-map-preview/` — превью стиля  
- `МосХаб/.rr-install/` — ключи RR (секреты)  

После старта в OnlyGo: только **код решения** + README/презентация, дробными коммитами. Правила ИИ остаются в `hotel/`.

---

## 1. Продуктовая рамка для задачи PT

### Проблема → гипотеза

**Проблема:** аналитику ИБ/DevSecOps по логам Runtime Radar тяжело за минуты собрать картину «что произошло внутри контейнера»: процессы разрознены, сеть отдельно, хронология в табличке.

**Гипотеза (проверяемая на демо):**  
Если в одном окне показать дерево процессов + сетевые связи + таймлайн с синхронизацией клика, то аналитик восстановит цепочку атаки (например reverse-shell) за < 60 секунд без ручного grep.

**Не питчить:** «мы сделали красивый UI».  
**Питчить:** «мы сокращаем время расследования инцидента в контейнере».

### Для кого

Целевой пользователь на питче: **SOC / container security analyst**.  
Боль: время от алерта RR до понимания process tree + сетевого C2.  
Ценность: один экран вместо трёх вкладок Runtime Radar.

### Метрики успеха на хакатоне

| Метрика | Цель на ночь |
|---------|----------------|
| Time-to-insight на демо | < 60 сек от открытия UI до «вот цепочка» |
| Покрытие ТЗ | tree + timeline + network в одном окне |
| Стабильность демо | fallback fixture если RR мёртв |
| История git | чистые conventional commits после старта |

### Структура питча (~3–5 мин, шаблон финалиста в OnlyGo)

1. **Крючок (15–20 сек):** «Контейнер уже скомпрометирован. В логах тысяча строк. Сколько минут до понимания?»  
2. **Проблема:** Runtime Radar даёт события, но не «картину в одном окне».  
3. **Продукт:** Activity Map — дерево / сеть / хронология, клик синхронизирует всё.  
4. **Демо (90–120 сек):** live или fixture reverse-shell.  
5. **Потенциал:** встраивание рядом с RR, replay, экспорт в SOC.  
6. **Команда / стек (20 сек).**  

Говорить коротко и по делу: один главный месседж — «это ускоряет расследование».

---

## 2. Критерии жюри → как закрываем

| Критерий | Вес | Как закрываем |
|----------|-----|----------------|
| Решение задания | 30% | Одно окно: хронология + tree процессов + сеть по данным RR |
| Качество | 25% | Чистый UI, контракт API, README, стабильный fallback |
| Инновационность | 20% | Replay цепочки / корреляция process↔net / «детектив»-режим |
| Презентация и команда | 15% | Сценарий питча + роли на демо |
| Потенциал | 10% | Roadmap: webhook в SIEM, multi-cluster, SDK детекторов |

---

## 3. Роли (4 человека) + работа с нейроагентами

Все четверо работают **с нейросетями**. Поэтому главное — не «у каждого своя ветка навсегда», а **границы зон** и **handoff с критериями приёмки**, чтобы агент P3 не начал «улучшать» BFF P2.

| Роль | Зона файлов (агенту можно трогать только это) | До старта | После старта |
|------|-----------------------------------------------|-----------|--------------|
| **P1 Platform** | `scripts/`, `deploy/`, `fixtures/` (сырые дампы RR) | Docker/kind/RR живы | gen-activity, port-forward, экспорт events |
| **P2 Backend** | `apps/bff/**`, `docs/CONTRACT.md` | Контракт в голове | `/api/activity`, normalize, fixture fallback |
| **P3 Frontend** | `apps/web/**`, `apps/web` styles | Preview стиля глазами | tree + network + timeline + sync |
| **P4 Pitch** | `docs/pitch/**`, `README.md`, презентация | Черновик **локально** | слайды, демо-скрипт, тайминг |

**Общее (никто не ломает без синка):** схема папок, OpenAPI/Pydantic-контракт; правила агентов — только в `hotel/AGENTS.md`.

Sync: каждые **60–90 мин**, 5 минут: что готово / что блочит / кого ждём по handoff.

### Как работать вместе, если у всех агенты

**1. Ownership по директориям (главное правило)**  
Path ownership — в `hotel/AGENTS.md`. Промпт агенту всегда начинать с:  
«Работай только в `apps/web/**`. Чужие папки не менять. Контракт API — read-only.»

**2. Контракт раньше кода**  
Первый час: заморозить `GET /api/activity` (shape из `stubs/CONTRACT.md`).  
FE и BE агенты стыкуются по JSON, а не «как получится». Breaking change контракта — только на sync и с апдейтом обеих сторон.

**3. Не «четыре бесконечные ветки», а короткие MR + один integrator**  
- Общая линия: `main` (или `develop`) после старта.  
- Короткоживущие ветки под задачу: `feat/bff-fixture`, `feat/web-tree` — мерж как только AC зелёные.  
- Один **integrator** (обычно P2 или капитан) принимает MR и гоняет smoke.  
Личные ветки «на всю ночь» с AI почти всегда расходятся и больно склеиваются.

**4. Стык ролей = код контракта, не бумажный ад**  
P2 с первого коммита BFF фиксирует **Pydantic-модели** → FastAPI отдаёт `/openapi.json`.  
P3 генерит TypeScript-типы из OpenAPI. Проверка стыка: `pytest` на схему + `curl`.  
Markdown handoff в `hotel/handoffs/` — опционально; в чате достаточно AC + ссылка на OpenAPI.

**5. ACID приёмки (локально в `hotel/skills/06-handoff-acid.md`)**  

| Буква | Смысл |
|-------|--------|
| **A**cceptance | явные критерии готовности |
| **C**ontract | OpenAPI/Pydantic не ломаем без синка |
| **I**solation | агент только в своей папке; корень — integrator |
| **D**emo-ready | curl или клик проверки |

**6. Чужое улучшать** только после явного handoff/AC или joint-сессии.

---

## 4. Правила агентов живут в `hotel/` — не в OnlyGo

Согласовано с риском «идеальный AGENTS.md через 5 минут после старта»:

- **`МосХаб/hotel/AGENTS.md`** + **`hotel/skills/`** + **`hotel/prompts/`** — на диске для Cursor/агентов.  
- В OnlyGo **не коммитим** AGENTS.md / skills / промпты.  
- При необходимости: `OnlyGo/.git/info/exclude` (локально), чтобы случайно не закоммитить.  
- Skills только в `hotel/skills/` (+ `hotel/SKILLS.md`); `claude-skills-main` агентам не подключать.

### Что коммитить в OnlyGo после старта (pacing)

1. `chore: scaffold apps/web and apps/bff`  
2. `feat: pydantic activity models + fixture endpoint`  
3. `feat: process tree with xyflow`  
4. `feat: timeline + network panels`  
5. `feat: selection sync + incident summary`  
6. `feat: live Runtime Radar client` (когда стабильно)  
7. `docs: README and demo notes`  

Не вываливать 40 файлов одним коммитом через 15 минут.

### Стек UI/API (зафиксировано в hotel)

- Web: **`@xyflow/react` + dagre** — не invent SVG graph math  
- BFF: FastAPI + Pydantic  
- Fixture-first с `X-Data-Source`  
- Корень репо: только integrator (P1/P2)

---

## 5. Архитектура решения

```
┌─────────────────┐     ┌──────────────────┐     ┌────────────────────────────┐
│ Runtime Radar   │────▶│ Activity BFF     │────▶│ Web: Activity Map          │
│ HTTP via :9000  │     │ Pydantic+OpenAPI │     │ xyflow tree | Network |    │
│ (or fixture)    │     │ X-Data-Source    │     │ Timeline | Incident Summary│
└─────────────────┘     └──────────────────┘     └────────────────────────────┘
```

**MVP UX (ТЗ):** выбор пода; дерево процессов; сеть; хронология; sync по клику.

**Инновация (20%):**  
1) Replay по таймлайну;  
2) **Incident Summary** — карточка для SOC (текст + copy) по выбранному узлу.

**Данные:** `fixtures/sample-activity.json`; live — HTTP API RR (см. § ниже).

### Как забираем события из живого Runtime Radar

Уже нащупан путь через **HTTP reverse-proxy** (`https://localhost:9000`), UI/API под admin-токеном (см. `.rr-install/ACCESS.md`, history/events в Runtime UI).  
План BFF: клиент к known REST endpoints public-api / history (уточнить по OpenAPI/Network tab в браузере в T0), нормализация в Pydantic.  
**Не** парсить stdout логов контейнера как основной путь.  
Если live мёртв — fixture, честно в UI/README.

---

## 6. Чеклист ДО старта (сегодня / до 7.11 вечера)

### Инфра (P1 + все)

- [ ] `docker info` ок  
- [ ] kind cluster `moshub` / `kubectl get nodes`  
- [ ] Runtime Radar pods Running; UI `https://localhost:9000` (admin / Password)  
- [ ] Clusters: kind-moshub Connected; Runtime → Events не пустые  
- [ ] Не включать `kernel-modules` на Docker Desktop kind  
- [ ] SSH к hub.mos.ru работает; OnlyGo clone ок  
- [ ] В OnlyGo видны 2 PDF + pptx  

### Команда

- [ ] Распределены роли P1–P4  
- [ ] Все прочитали критерии и карточку ТЗ  
- [ ] Все видели локальный preview стиля (`activity-map-preview` :5177)  
- [ ] Договорились: **ноль пушей кода в OnlyGo до старта**  
- [ ] Черновик питча у P4 локально (Google Doc / файл не в OnlyGo)  

### Запрещённый «предкоммит» / не светить ИИ-инструкции

- [ ] Нет `apps/`, решения, `AGENTS.md`, `skills/` в OnlyGo до старта  
- [ ] После старта тоже не коммитить AGENTS/skills — они в `hotel/`  
- [ ] Секреты `.rr-install` не в git  

---

## 7. План ночи (после официального старта)

### T0 (первые 20–30 мин) — Kickoff на репо

1. Все pull OnlyGo.  
2. **Не** пушить AGENTS/skills — открыть `МосХаб/hotel/AGENTS.md` локально у агентов.  
3. P2: scaffold BFF + Pydantic models (первый нормальный коммит кода).  
4. Standup: зоны + кто integrator.

### T1 (час 1–2) — Scaffold / fixture-first

- P2: fixture endpoint + OpenAPI + pytest схемы  
- P3: Vite React + xyflow + токены стиля (по частям из kit)  
- P1: scripts port-forward / gen-activity  
- P4: презентация в OnlyGo **после старта**, короткими docs-коммитами

### T2–T4 — Core MVP

- P3: Process tree SVG + timeline + network panels + selection sync  
- P2: normalize live RR events (если API доступен) + merge с fixture  
- P1: стабильный demo-pod сценарий  
- P4: минутный сценарий демо, скринкаст-страховка  

### T4–T6 — Polish + innovation

- Replay button  
- **Incident Summary** (SOC-карточка + copy)  
- Фильтры по типу события  
- Seamless `X-Data-Source: fixture|live`  
- README: 3 команды  

### T6–T7 — Freeze

- E2E с нуля на одном ноутбуке  
- Public visibility OnlyGo  
- Репетиция питча 2×  
- CHANGELOG / список коммитов для жюри  

---

## 8. Что делать с kit / hotel / preview

- `hotel/` — промпты и AGENTS для агентов на диске.  
- `hackathon-kit/` — стиль + fixture + план.  
- Писать код в OnlyGo **с нуля после старта**, ориентируясь на kit, без dump и без коммита hotel.  
- Pacing: естественные инженерные коммиты (§4).  

---

## 9. Риски

| Риск | Митигация |
|------|-----------|
| Жюри видит предкоммиты | Не пушить код до старта |
| RR monitor падает на kind | Fixture + честный README |
| Docker down ночью | P1 чеклист; второй ноут |
| Рассинхрон FE/BE | Контракт с T0, breaking freeze после T3 |
| Питч без демо | Скринкаст + fixture mode |

---

## 10. Definition of Done

- [ ] В OnlyGo после старта есть история коммитов решения  
- [ ] UI: хронология + дерево процессов + сеть в одном окне  
- [ ] Данные из RR и/или явный fixture fallback  
- [ ] README + презентация  
- [ ] 3-мин демо отрепетировано  
- [ ] Visibility Public (как требуют правила сдачи)  
- [ ] Секретов в репо нет  

---

## 11. Быстрые ссылки локально

- План: этот файл  
- **Hotel (агенты):** `../hotel/` — AGENTS, skills, prompts  
- Стиль / fixture: `design/`, `fixtures/`  
- Контракт-черновик: `stubs/CONTRACT.md` (в коде — Pydantic)  
- Preview: `../activity-map-preview` → http://127.0.0.1:5177/  
- RR доступ: `../.rr-install/ACCESS.md`  
- Skills: `../hotel/SKILLS.md` + `../hotel/skills/` (claude-skills-main не использовать)  
