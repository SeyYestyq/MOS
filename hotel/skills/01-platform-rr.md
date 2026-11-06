# Platform / Runtime Radar (P1)

Всё нужное здесь. Не открывать целиком чужие skill-репо.

## Зона

Писать: `scripts/**`, `deploy/**`, `fixtures/raw/**`.  
Не трогать: `apps/**`, корень репо без роли integrator.

## Задачи ночи

1. kind + RR живы (`kubectl`, port-forward `:9000`).
2. Источник событий: history/sources/allowlist + demo-pod при пустых Events.
3. Сдампить сырой ответ/экспорт в `fixtures/raw/` для P2.
4. Скрипт «поднять демо-активность» (curl/shell), один README-шаг для питча.
5. Если RR падает — triage ≤10 мин → fixture (`07-focused-fix.md`); не чинить kind всю ночь.

## Паттерны (сжато)

- Идемпотентные скрипты: повторный запуск безопасен.
- Секреты только в `.rr-install` / env (`08-git-secrets.md`).
- Health: `kubectl -n runtime-radar get pods`, UI `:9000`, Events не пустые.
- Rollback: fixture dump; live — бонус.
- Docker/compose советы при сборке демо: `_borghei/docker-development.md` (точечно).

## Не делать

Terraform, blue/green, CI generators, cloud architect skills.
