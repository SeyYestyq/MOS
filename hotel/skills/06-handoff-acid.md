# 06 — Handoff ACID (для людей и нейроагентов)

Только локально в `hotel/skills/`. В OnlyGo **не** коммитить.

## Path isolation

Агент меняет только пути своей роли (см. AGENTS.md). Чужие директории — read-only.

## Handoff card template

Класть в `docs/handoffs/YYYYMMDD-from-to-topic.md`:

```md
# Handoff: <from> → <to>
Goal: ...
## Acceptance Criteria
- [ ] ...
## How to verify
...
## Out of scope
...
```

## ACID

- **A**cceptance — чеклист `- [ ]`, без него MR не принимаем  
- **C**ontract — не ломать `docs/CONTRACT.md` без синка команды  
- **I**solation — не писать вне своей зоны  
- **D**emo-ready — есть команда/клик проверки  

## Definition of Done (фича)

AC зелёные + conventional commit + зона не расползлась + verify пройден.
