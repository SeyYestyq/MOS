# Docs / CHANGELOG / BASE (роль C)

Паттерн как у длинных проектов (changelog + короткие отчёты), урезанно под 8 часов.  
Шаблоны: `hackathon-kit/stubs/`.

## Что лежит в OnlyGo (после старта)

| Файл | Зачем |
|------|--------|
| `CHANGELOG.md` | Журнал этапов ночи для жюри / истории |
| `BASE.md` | **Обязательный базовый файл** из PDF PT (+ ИИ) |
| `docs/README.md` | Оглавление |
| `docs/DEMO.md` | Сценарий демо < 60 с |
| `docs/architecture.md` | RR → BFF → Web в 10–20 строк |
| `README.md` | 3 команды запуска |
| `reports/SUMMARY_*.md` | Опционально после крупного куска |

## Правило CHANGELOG

После каждого **осмысленного** этапа (не каждый микрокоммит) добавить блок:

```markdown
## YYYY-MM-DD HH:MM — кратко что

- **Кто:** A | B | C
- **Сделано:** …
- **Проверка:** curl / клик / команда
```

Примеры этапов: fixture API жив; tree на xyflow; sync панелей; Incident Summary; live RR; README готов.

## Правило BASE.md (сдача)

Перед freeze заполнить полностью:

1. Название, команда, задача PT Activity Map  
2. Что умеет UI (tree / network / timeline / sync / summary)  
3. Стек и запуск  
4. **ИИ-инструменты и этапы** (Cursor и т.д. — scaffold / UI / BFF / docs)  
5. Решения приняты и объяснимы командой  
6. ИБ-проверка: нет секретов в репо, токены из env, код от ИИ просмотрен  

Не копировать hotel/AGENTS в BASE. Не писать пароли RR.

## Conventional commits

`feat(web):` · `feat(bff):` · `docs:` · `fix:` · `chore:`  
Scopes: `web` | `bff` | `scripts` | `docs` | `fixtures`

## Не делать

Тащить gotender-style тонны reports; дублировать ТЗ целиком; коммитить skills.
