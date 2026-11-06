# Freeze review (капитан / любой перед сдачей)

Всё нужное здесь. Не гонять code-reviewer scripts из claude-skills.

## За 30–45 мин до freeze

- [ ] Демо проходит на fixture без кластера  
- [ ] Live RR — если работает; иначе честно fixture в UI  
- [ ] Нет секретов/токенов в OnlyGo  
- [ ] Нет `AGENTS.md`, `skills/`, `hotel/` в истории коммитов OnlyGo  
- [ ] Path ownership не сломан (web не в bff и наоборот)  
- [ ] OpenAPI и UI согласованы (поля дерева/таймлайна есть)  
- [ ] Incident Summary копируется  
- [ ] README: как запустить за < 3 команд  
- [ ] Коммиты выглядят как этапы, не один dump на 40 файлов в T+15м  

## Красные флаги в diff

Hardcoded password/token, `console.log` с dump всего ответа RR, правка чужой зоны, breaking rename полей без синка.
