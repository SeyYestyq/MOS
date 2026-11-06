# hotel/ — только локально, никогда в OnlyGo

Правила для нейроагентов. Агенты читают с диска. В Mos.Хаб — только код решения.

## Содержимое

| Путь | Зачем |
|------|--------|
| `AGENTS.md` | Закон: зоны, стек, pacing |
| `SKILLS.md` | Карта skills (всё внутри hotel) |
| `skills/` | Самодостаточные инструкции по ролям |
| `prompts/` | Брифы в буфер агенту |
| `handoffs/` | Опционально; предпочтительно OpenAPI |

`claude-skills-main` рядом можно оставить на диске как архив — **агентам его не подключать**.

## Exclude в OnlyGo (локально)

```bash
echo 'AGENTS.md' >> .git/info/exclude
echo 'skills/' >> .git/info/exclude
echo '.hotel' >> .git/info/exclude
```
