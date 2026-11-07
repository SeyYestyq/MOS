# Выжимка: «Материалы по Runtime Radar.docx»

Источник: `../Материалы по Runtime Radar.docx` (локально).

## Стек решения (что говорит материал)

- **UI:** React + TypeScript + Vite + **Zustand** + **React Flow + dagre** + **vis-timeline**  
- **Интеграция:** public-api / reverse-proxy + **JSON-выгрузка** Events (обязательный fallback); опционально history-api / webhook  
- **Сам Runtime Radar** написан на **Go** (микросервисы) — это продукт PT, **не** требование писать весь хакатонный UI на Go  
- eBPF/Tetragon/BTF — инфраструктура сбора, не язык вашего фронта

## Критично: дерево процессов

**Не** связывать только по `pid`/`ppid` (коллизии в контейнерах).  
Связь узлов: **`process.exec_id` → `parent.exec_id`** (DAG).  
Дополнительно: `start_time`, `time`, `in_init_tree`, `binary`, `arguments`, `cwd`.

Сеть: из `PROCESS_KPROBE` — `saddr/sport/daddr/dport/family`, привязка к `exec_id` процесса.

## Тестовая активность (Platform)

Namespace `hackathon`, nginx + alpine, curl/wget, затем `nc -l -p 8888` + `ls` (триггер детектора / MITRE TA0011 T1571).  
Sources + Allow filter + **All**. UI `:32000` или port-forward `:9000`.

## Инновация / питч

Brushing & Linking (клик таймлайн ↔ дерево ↔ сеть); подсветка пути атаки + бейджи MITRE; storytelling < 15–60 с.

## Потенциал

UI как модуль, совместимый по идее с `radar-ui` / public-api / history-api.
