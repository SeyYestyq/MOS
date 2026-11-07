/**
 * STUB — карточка для SOC + copy (инновация питча).
 * Текст по шаблону hotel/skills/09-soc-demo-language.md
 * + Инновация: генератор YAML политики TracingPolicy для Tetragon.
 */
import { useState } from "react";
import type { ProcessNode, NetEdge, ActivityPayload } from "../types";

export type IncidentSummaryProps = {
  payload: ActivityPayload;
  selected: ProcessNode | null;
  relatedNet: NetEdge[];
  source?: "fixture" | "live";
};

export function buildIncidentText(
  payload: ActivityPayload,
  selected: ProcessNode | null,
  relatedNet: NetEdge[],
  source: "fixture" | "live" = "fixture",
): string {
  if (!selected) return "Выберите процесс в дереве или на таймлайне для экспресс-анализа.";
  
  // Связь по execId в приоритете, с fallback на pid
  const parent = selected.parentExecId
    ? payload.processes.find((p) => p.execId === selected.parentExecId)
    : payload.processes.find((p) => p.pid === selected.ppid);

  const net = relatedNet[0];
  const endpoint = net
    ? `${net.dst || net.daddr}${net.port != null ? `:${net.port}` : ""}`
    : "отсутствует";
  
  const when = selected.lastSeen || selected.firstSeen || selected.startTime || "н/д";
  const mitreText = selected.mitre && selected.mitre.length > 0
    ? `\n[ВНИМАНИЕ] Сработал детектор угроз: ${selected.mitre.join(", ")}`
    : "";

  return [
    `Инцидент в поде ${payload.pod} (namespace: ${payload.namespace}):`,
    `Родительский процесс: ${parent?.binary ?? "init/unknown"} (PID: ${parent?.pid ?? selected.ppid ?? "n/a"})`,
    `Порожденный процесс: ${selected.binary} (PID: ${selected.pid}, exec_id: ${selected.execId})`,
    `Командная строка: ${selected.arguments ?? "n/a"}`,
    `Сетевая активность: ${endpoint} (протокол: ${net?.proto ?? "n/a"}) в ${when} UTC.${mitreText}`,
    `Источник телеметрии: ${source === "live" ? "Runtime Radar (Live Cluster)" : "Эталонный дамп (Fixture)"}.`,
  ].join("\n");
}

export function generateTracingPolicyYaml(selected: ProcessNode | null, relatedNet: NetEdge[]): string {
  if (!selected) return "";
  const net = relatedNet[0];
  const port = net?.port ?? net?.dport ?? 8888;
  const binary = selected.binary || "/usr/bin/nc";

  return `apiVersion: cilium.io/v1alpha1
kind: TracingPolicy
metadata:
  name: block-${selected.binary.split("/").pop() || "threat"}-policy
spec:
  kprobes:
    - call: "inet_csk_listen_start"
      syscall: false
      args:
        - index: 0
          type: "sock"
      selectors:
        - matchArgs:
            - index: 0
              operator: "Equal"
              values: ["${port}"]
          matchActions:
            - action: Sigkill
    - call: "sys_enter_execve"
      syscall: true
      args:
        - index: 0
          type: "string"
      selectors:
        - matchArgs:
            - index: 0
              operator: "Prefix"
              values: ["${binary}"]
          matchActions:
            - action: Sigkill`;
}

export function IncidentSummary({
  payload,
  selected,
  relatedNet,
  source = "fixture",
}: IncidentSummaryProps) {
  const [showPolicy, setShowPolicy] = useState(false);
  const text = buildIncidentText(payload, selected, relatedNet, source);
  const policyYaml = generateTracingPolicyYaml(selected, relatedNet);

  return (
    <div className="admin-surface liquid-metal-frame liquid-metal-frame--soft">
      <div className="admin-surface-head">
        <h3>Context & Response Drawer</h3>
        <span className="badge-soc">SOC / SecOps</span>
      </div>

      <div style={{ display: "flex", gap: "8px", margin: "8px 0" }}>
        <button
          type="button"
          className={`btn-tab ${!showPolicy ? "active" : ""}`}
          onClick={() => setShowPolicy(false)}
        >
          Incident Summary
        </button>
        <button
          type="button"
          className={`btn-tab ${showPolicy ? "active" : ""}`}
          onClick={() => setShowPolicy(true)}
        >
          TracingPolicy (YAML)
        </button>
      </div>

      {!showPolicy ? (
        <>
          <pre className="incident-summary">{text}</pre>
          <div className="incident-summary__actions">
            <button
              type="button"
              className="btn-copy"
              onClick={() => navigator.clipboard.writeText(text)}
            >
              Скопировать отчет для руководства
            </button>
          </div>
        </>
      ) : (
        <>
          <pre className="incident-summary" style={{ fontSize: "11px", color: "var(--neon-green, #10b981)" }}>
            {policyYaml || "# Выберите инцидент для генерации правила"}
          </pre>
          <div className="incident-summary__actions">
            <button
              type="button"
              className="btn-copy"
              disabled={!policyYaml}
              onClick={() => navigator.clipboard.writeText(policyYaml)}
            >
              Скопировать TracingPolicy YAML
            </button>
          </div>
        </>
      )}
    </div>
  );
}
