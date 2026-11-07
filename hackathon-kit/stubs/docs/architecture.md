# Архитектура (кратко)

```
Runtime Radar (HTTPS :9000)
        │  HTTP API / dump
        ▼
   Activity BFF (Go)
        │  GET /api/activity
        │  X-Data-Source: fixture | live
        ▼
   Web Activity Map (Vite/React)
        ├── Process tree (@xyflow/react + dagre, exec_id)
        ├── Network panel
        ├── Timeline (vis-timeline)
        └── Incident Summary
```

Контракт: Go structs ↔ `types.ts`. Дерево по `exec_id` / `parent_exec_id`.  
Fixture-first: демо не зависит от живого кластера.
