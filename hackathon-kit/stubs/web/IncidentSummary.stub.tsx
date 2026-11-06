/**
 * STUB — карточка для SOC + copy (инновация питча).
 * Текст по шаблону hotel/skills/09-soc-demo-language.md
 */
import type { ProcessNode, NetEdge, ActivityPayload } from "../types";

export type IncidentSummaryProps = {
  payload: ActivityPayload;
  selected: ProcessNode | null;
  relatedNet: NetEdge[];
};

export function buildIncidentText(
  payload: ActivityPayload,
  selected: ProcessNode | null,
  relatedNet: NetEdge[],
  source: "fixture" | "live" = "fixture",
): string {
  if (!selected) return "Выберите процесс в дереве.";
  const parent = payload.processes.find((p) => p.pid === selected.ppid);
  const net = relatedNet[0];
  const endpoint = net
    ? `${net.dst}${net.port != null ? `:${net.port}` : ""}`
    : "n/a";
  const when = selected.lastSeen || selected.firstSeen || "n/a";
  return [
    `Аномалия в поде ${payload.pod} (ns ${payload.namespace}):`,
    `процесс ${parent?.binary ?? "?"} (PID ${selected.ppid}) породил ${selected.binary} (PID ${selected.pid})`,
    `с внешним соединением ${endpoint} в ${when}.`,
    `Источник данных: ${source}.`,
  ].join("\n");
}

export function IncidentSummary({
  payload,
  selected,
  relatedNet,
}: IncidentSummaryProps) {
  const text = buildIncidentText(payload, selected, relatedNet);
  return (
    <div className="admin-surface liquid-metal-frame liquid-metal-frame--soft">
      <div className="admin-surface-head">
        <h3>Incident Summary</h3>
        <span>SOC</span>
      </div>
      <pre className="incident-summary">{text}</pre>
      <div className="incident-summary__actions">
        <button
          type="button"
          className="btn-copy"
          onClick={() => navigator.clipboard.writeText(text)}
        >
          Copy to ticket
        </button>
      </div>
    </div>
  );
}
