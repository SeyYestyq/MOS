# Frontend Activity Map (P3)

Всё нужное здесь. Не открывать `claude-skills-main` / senior-frontend.

## Зона

Только `apps/web/**`.

## Стек (фиксировано)

- Vite + React + TypeScript
- **`@xyflow/react`** + **`@dagrejs/dagre`** для дерева процессов
- Типы из OpenAPI BFF (сгенерировать или скопировать вручную один раз)
- Стиль: токены из `hackathon-kit/design/`
- **Не** рисовать связи ppid→pid на чистом SVG/Canvas/D3 с нуля

## UI блоки (обязательные)

1. Process tree (кастомные ноды процессов)
2. Network panel / edges
3. Timeline
4. Selection sync (клик в любом блоке подсвечивает связанные)
5. **Incident Summary** — текст для SOC по выбранному узлу + copy в буфер

## Порядок коммитов (pacing)

1. Vite scaffold + layout shell + токены  
2. Fetch `/api/activity` + типы  
3. Дерево на React Flow + dagre  
4. Network + timeline  
5. Sync кликов  
6. Incident Summary + polish  

## Паттерны UI

- Один shared selection state (id процесса / event ts).
- Fixture работает без живого BFF через mock или заголовок.
- Анимации умеренно: highlight + pan-to-node, без шума.
- a11y минимум: фокус, контраст, кнопки с label.

## Не делать

Next.js App Router, Tailwind-only редизайн с нуля, Storybook, auth features scaffolder.
