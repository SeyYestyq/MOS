# Freeze review (капитан / любой перед сдачей)

Всё нужное здесь. Не гонять code-reviewer scripts из claude-skills.

## За 30–45 мин до freeze

- [ ] Демо проходит на fixture без кластера  
- [ ] Live RR — если работает; иначе честно fixture в UI  
- [ ] Нет секретов/токенов в OnlyGo  
- [ ] Нет `AGENTS.md`, `skills/`, `hotel/` в истории коммитов OnlyGo  
- [ ] Path ownership Frontend/Backend/Platform не сломан  

- [ ] OpenAPI и UI согласованы  
- [ ] Incident Summary копируется  
- [ ] README + **CHANGELOG.md** + **BASE.md** (блок ИИ заполнен)  
- [ ] Чеклист ИБ `11-secure-solution.md` зелёный  
- [ ] Коммиты = этапы, не mega-dump в T+15м  



## Красные флаги в diff

Hardcoded password/token, `console.log` с dump всего ответа RR, правка чужой зоны, breaking rename полей без синка, `Co-authored-by: Cursor`.

Секреты/коммиты: `08-git-secrets.md`. Точечно: `_borghei/code-reviewer.md`.
