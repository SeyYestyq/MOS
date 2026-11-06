# SOC / demo language — из borghei threat-detection + senior-secops + roadmap

## Словарь для UI и питча

Говорить как аналитик, не как дизайнер:

- process tree / parent→child / spawn
- reverse shell / C2 / external egress
- timeline / time-to-insight
- Incident Summary → тикет SOC

## Шаблон Incident Summary (генерация в UI)

```
Аномалия в поде {pod} (ns {namespace}):
процесс {parent} (PID {ppid}) породил {child} (PID {pid})
с внешним соединением {ip}:{port} в {time} UTC.
Источник данных: {fixture|live}.
```

Кнопка Copy → буфер.

## Демо-скрипт жюри (< 60 с)

1. Открыть Activity Map (fixture ok)  
2. Клик suspect `/bin/sh` или python payload  
3. Подсветка сети `:4444` + событие на таймлайне  
4. Incident Summary → Copy  
5. Одна фраза: «вместо трёх вкладок RR — одна картина»

## Now / Next / Later (питч)

Не обещать даты. Confidence: Now = сделано, Next = логично, Later = горизонт.

Детали: `_borghei/threat-detection.md`, `_borghei/senior-secops.md`, `_borghei/roadmap-communicator.md`
