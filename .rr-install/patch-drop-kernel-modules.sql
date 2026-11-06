-- Kind/linuxkit: drop kernel-modules TracingPolicy (free_module kprobe missing).
-- Apply against DB runtime-radar, then restart runtime-monitor DaemonSet pod.
UPDATE runtime_monitor_configs
SET config = config #- '{tracing_policies,kernel-modules}';
SELECT jsonb_object_keys(config->'tracing_policies') AS policies FROM runtime_monitor_configs;
