/**
 * STUB — реализовать после старта в apps/web.
 * Движок: @xyflow/react + @dagrejs/dagre (иерархическая укладка TB)
 * Паттерн: ego-graph фокус вокруг выбранного процесса
 * Стили: import styles/activity-map.css (tokens + topology + xyflow-theme)
 * Типы: stubs/types.ts → apps/web/src/types.ts
 */
import type { ProcessNode } from "../types";

export type ProcessTreeBoardProps = {
  processes: ProcessNode[];
  selectedId: string | null;
  onSelect: (node: ProcessNode) => void;
  attackPathActive?: boolean;
};

export function ProcessTreeBoard(_props: ProcessTreeBoardProps) {
  return (
    <div className="admin-surface liquid-metal-frame liquid-metal-frame--soft app-shell__tree">
      <div className="admin-surface-head">
        <h3>Process Tree (Dagre DAG)</h3>
        <span className="badge-status">@xyflow/react + Dagre</span>
      </div>
      <p style={{ color: "var(--text-muted)", fontSize: "13px", padding: "12px" }}>
        Иерархический граф процессов, построенный строго по <code>child.parentExecId === parent.execId</code>.
        Кастомные ноды <code>.process-node</code> с отображением командной строки, флагов MITRE и поддержкой ego-graph фокусировки.
      </p>
    </div>
  );
}
