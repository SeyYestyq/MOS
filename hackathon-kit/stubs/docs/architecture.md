# Архитектура (кратко)

```
Runtime Radar (HTTPS :9000)
        │  HTTP API / dump
        ▼
   Activity BFF (FastAPI + Pydantic)
        │  GET /api/activity
        │  X-Data-Source: fixture | live
        ▼
   Web Activity Map (Vite/React)
        ├── Process tree (@xyflow/react + dagre)
        ├── Network panel
        ├── Timeline
        └── Incident Summary
```

Контракт: модели Pydantic → `/openapi.json` → типы фронта.  
Fixture-first: демо не зависит от живого кластера.
