# Pitch / Demo (роль C)

Всё нужное здесь. Не открывать roadmap-communicator целиком.

## Зона

`docs/pitch/**`, `README.md`, слайды, совместно с `10-docs-changelog.md` (`BASE.md`, `DEMO.md`).  
Не править `apps/**` без joint.

## Структура питча (5–7 мин)

1. **Проблема SOC** — 3 вкладки, долгое расследование reverse-shell  
2. **Метрика** — Time-to-Insight < 60 секунд в одном окне  
3. **Живое демо** — дерево → сеть → таймлайн → Incident Summary → copy  
4. **Как устроено** — RR / fixture, BFF контракт, карта  
5. **Потенциал** — Now / Next / Later (без ложных дедлайнов)

## Now / Next / Later

- **Now:** MVP карта + sync + summary + fixture fallback  
- **Next:** live RR стабильнее, фильтры, экспорт в тикет  
- **Later:** корреляции между подами, playbooks  

## README (коротко)

Проблема → запуск (web + bff + fixture) → что показать жюри → ограничения.

## Чеклист качества текста

- [ ] Язык аналитика/SOC, не «наш крутой граф»  
- [ ] Одна метрика успеха  
- [ ] Риски (kind/eBPF) названы, fixture как страховка  
- [ ] Нет секретов RR в слайдах/README  

Язык SOC + секундный демо-скрипт: **`09-soc-demo-language.md`**.

## Не делать

Полный OKR cascade, board-deck на 20 слайдов, marketing landing generator.
