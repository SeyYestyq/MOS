# Runtime Radar — доступ (kind-moshub)

## Установка
- Chart: `oci://ghcr.io/runtime-radar/runtime-radar:v0.2.0`
- Namespace: `runtime-radar`
- Release: `runtime-radar`
- Values override: `.rr-install/kind-values.yaml` (теги образов v0.2.0, 1 реплика)

## Логин UI
- username: `admin`
- password: `Password`

## Открыть UI (рекомендуется port-forward)
```bash
kubectl --context kind-moshub -n runtime-radar port-forward svc/reverse-proxy 9000:9000
```
Открыть в браузере: https://localhost:9000  
(самоподписанный сертификат — принять исключение)

Страница Runtime: https://localhost:9000/runtime (или `/runtime/events`, `/runtime/settings`).  
Отдельный query-param с cluster id для Runtime **не нужен** — кластер выбирается через Switch/Clusters.

Альтернатива (NodePort 32000): на Docker Desktop kind обычно **не** проброшен наружу без extraPortMappings.

## Kind / Docker Desktop workaround (runtime-monitor)

На linuxkit (Docker Desktop kind) kprobe `free_module` отсутствует → runtime-monitor падает при загрузке TracingPolicy `kernel-modules`, endpoint `/api/v1/config/runtime-monitor` отдаёт 502, UI Runtime показывает «404».

Патч (уже применён на текущем кластере): удалить политику из Postgres и перезапустить DaemonSet:

```bash
# credentials из secret/postgresql
kubectl --context kind-moshub -n runtime-radar exec -i postgresql-0 -- \
  env PGPASSWORD='…' psql -U runtime-radar -d runtime-radar \
  -c "UPDATE runtime_monitor_configs SET config = config #- '{tracing_policies,kernel-modules}';"
kubectl --context kind-moshub -n runtime-radar delete pod -l app.kubernetes.io/name=runtime-monitor
# startupProbe ~70s — подождать Ready 2/2
```

Не включайте снова source «kernel-modules» в UI на kind — sidecar снова уйдёт в CrashLoop.  
Полный eBPF на kind/linuxkit ненадёжен; для демо событий лучше mock/Linux VM.
