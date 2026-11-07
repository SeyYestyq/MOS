# hotel/ — только локально, никогда в OnlyGo

Правила для нейроагентов (команда **A / B / C**). В Mos.Хаб — код решения + docs/BASE/CHANGELOG.

## Содержимое

| Путь | Зачем |
|------|--------|
| `AGENTS.md` | Зоны A/B/C, стек, pacing |
| `SKILLS.md` | Карта skills 01–10 |
| `skills/` | Ролевые инструкции + `_borghei/` |
| `prompts/` | Брифы агенту |
| `handoffs/` | Опционально |

## Exclude в OnlyGo (локально)

```bash
echo 'AGENTS.md' >> .git/info/exclude
echo 'skills/' >> .git/info/exclude
echo '.hotel' >> .git/info/exclude
```
