# Platform / Runtime Radar (P1)

Всё нужное здесь. Не открывать `claude-skills-main`.

## Зона

Писать: `scripts/**`, `deploy/**`, `fixtures/raw/**`.  
Не трогать: `apps/**`, корень репо без роли integrator.

## Задачи ночи

1. kind + RR живы (`kubectl`, port-forward `:9000`).
2. Источник событий: history/sources/allowlist + demo-pod при пустых Events.
3. Сдампить сырой ответ/экспорт в `fixtures/raw/` для P2.
4. Скрипт «поднять демо-активность» (curl/shell), один README-шаг для питча.
5. Если RR падает — triage 10 мин, потом команда на fixture; не чинить кластер всю ночь.

## Паттерны (сжато)

- Идемпотентные скрипты: повторный запуск безопасен.
- Секреты только в `.rr-install` / env, не в OnlyGo.
- Health: `kubectl -n runtime-radar get pods`, UI `:9000`, Events не пустые.
- Rollback: вернуться на fixture dump; live — бонус.

## Не делать

Terraform, blue/green, CI pipeline generators, cloud architect skills.
