# Platform / Runtime Radar

Инженерная роль (не «человек на презентацию»).

## Зона

`scripts/**`, `deploy/**`, `fixtures/raw/**`.  
Не трогать `apps/**` без просьбы Frontend/Backend.

## Задачи

1. RR жив: pods, port-forward `:9000`.  
2. Sources; Allow filter ns; сохранение **All**.  
3. Тестовая активность: ns `hackathon`, nginx+alpine, `ls`/`cat`/`curl`/`wget`, затем `nc -l -p 8888` + `ls` (детектор/MITRE). См. `hackathon-kit/refs/RR_MATERIALS_NOTES.md`.  

4. Events не пустые.  
5. Dump → `fixtures/raw/` для Backend.  
6. RR упал → ≤10 мин triage → fixture (`07-focused-fix.md`).  
7. Подсказать Backend endpoints live API (DevTools Network).

## Не делать

Сидеть ночь на слайдах вместо инфра-задач; писать React/Pydantic за других.
