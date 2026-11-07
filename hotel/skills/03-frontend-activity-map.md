# Frontend Activity Map (роль A)

Всё нужное здесь. Не открывать чужие репозитории за стилем.

## Зона

Только `apps/web/**`.

## Стек (фиксировано)

- Vite + React + TypeScript
- **`@xyflow/react`** + **`@dagrejs/dagre`**
- Типы: скопировать `hackathon-kit/stubs/types.ts`
- Стиль: **только** `hackathon-kit/design/` → один импорт `activity-map.css`
- Fixture через BFF или `public/sample-activity.json`

## T0 — скопировать стили (не изобретать)

```text
hackathon-kit/design/tokens.css
hackathon-kit/design/topology.css
hackathon-kit/design/xyflow-theme.css
hackathon-kit/design/activity-map.css   ← import в main.tsx
hackathon-kit/design/STYLE_GUIDE.md     ← read-only шпаргалка
```

Классы панелей уже есть: `.admin-surface`, `.liquid-metal-frame`, `.process-node`, `.incident-summary`, `.app-shell`.

Превью: `activity-map-preview` — смотреть; в OnlyGo не dump-ить.

## UI блоки

1. Process tree (кастомные ноды `.process-node`, edges parent→child)
2. Network panel
3. Timeline
4. Selection sync
5. **Incident Summary** + copy

## Порядок коммитов

1. scaffold + `activity-map.css` + shell layout  
2. fetch/types + fixture  
3. React Flow + dagre  
4. network + timeline  
5. sync  
6. Incident Summary  

## Не делать

Чистый SVG/Canvas/D3 граф с нуля; Next.js; редизайн токенов; искать CSS вне `hackathon-kit/design/`.
