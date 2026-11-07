# Безопасность решения (все роли) — хакатон PT

Организаторы: решение должно быть **безопасным**; ИИ-код по ИБ **проверяет команда** (см. PDF / `BASE.md`).

Это не CTF и не полный SOC 2. Минимум для сдачи:

## Must (перед каждым push в OnlyGo)

- [ ] Нет паролей / токенов RR / encryption keys в git (`creds.txt`, `keys.txt`, живой `.env`)
- [ ] RR credentials только из env (`RR_PASSWORD`, `RR_TOKEN`) — шаблон `stubs/bff/env.example`
- [ ] Frontend не печатает секреты в UI / `localStorage` / URL
- [ ] `console.log` не сваливает полный ответ RR с токенами
- [ ] CORS: только origin Vite (не `*` с credentials в проде-демо ок аккуратно)
- [ ] Данные событий из RR/fixture — **текст как data**, не `dangerouslySetInnerHTML` без санитайза (XSS)
- [ ] Нет `eval` / shell из пользовательского ввода
- [ ] Зависимости только из npm/pypi по необходимости; не коммитить `node_modules`

## Backend

- Явные Go structs + json tags; не отдавать сырой dump RR клиенту
- Ошибки RR → fallback fixture, без утечки stack trace с секретами клиенту
- Не проксировать произвольные URL (SSRF): только фиксированный `RR_BASE_URL`

## Frontend

- Команды/cmdline процессов показывать как текст (escape по умолчанию в React)
- Copy Incident Summary — только собранная строка, без сырого dump всего API

## Platform

- Не класть admin Password в README/слайды/BASE (достаточно «см. локальный env»)
- Dump в `fixtures/raw/` без cookies/Authorization headers

## Freeze + BASE.md

- Пройти этот чеклист
- В `BASE.md` § ИБ: «код от ИИ просмотрен, секретов в репо нет»
- Skill детали (опционально): `_borghei/env-secrets-manager.md`, `_borghei/security-audit.md`, `_borghei/senior-secops.md`

## Не тащить на ночь

Полный PCI/HIPAA, Vault rotation, cloud-security architect, CTF-эксплойты.
