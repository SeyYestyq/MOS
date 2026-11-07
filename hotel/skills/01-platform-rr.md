# Platform / Runtime Radar (роль C)

Всё нужное здесь. Не открывать целиком чужие skill-репо.

## Зона

Писать: `scripts/**`, `deploy/**`, `fixtures/raw/**`.  
Docs/CHANGELOG/BASE — см. `10-docs-changelog.md`.  
Не трогать: `apps/**` без просьбы A/B.

## Задачи ночи

1. kind + RR живы (`kubectl`, port-forward `:9000`).
2. Sources; Allow filter на ns (напр. `hackathon`); сохранение **All**; Apply.
3. Тестовая активность: pods + `kubectl exec` (`ls`/`cat`/`sleep`/`sh -c`) + `curl`/`wget` с паузами.
4. Проверка: Runtime → Events не пустые.
5. Dump → `fixtures/raw/` для **B**.
6. Если RR падает — triage ≤10 мин → fixture (`07-focused-fix.md`).

## Паттерны

- Идемпотентные скрипты; секреты только локально (`08-git-secrets.md`).
- Health: pods Ready, UI `:9000`, Events видны.
- Rollback: fixture; live — бонус.

## Не делать

Terraform, blue/green, CI generators; писать React/BFF за A/B.
