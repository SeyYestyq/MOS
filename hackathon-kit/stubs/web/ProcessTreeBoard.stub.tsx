/**
 * STUB — реализовать после старта в apps/web.
 * Движок: @xyflow/react + @dagrejs/dagre
 * Стили: import styles/activity-map.css (tokens + topology + xyflow-theme)
 * Типы: stubs/types.ts → apps/web/src/types.ts
 */
import type { ProcessNode } from "../types";

export type ProcessTreeBoardProps = {
  processes: ProcessNode[];
  selectedId: string | null;
  onSelect: (node: ProcessNode) => void;
};

export function ProcessTreeBoard(_props: ProcessTreeBoardProps) {
  return (
    <div className="admin-surface liquid-metal-frame liquid-metal-frame--soft app-shell__tree">
      <div className="admin-surface-head">
        <h3>Process tree</h3>
        <span>TODO · React Flow</span>
      </div>
      <p style={{ color: "var(--text-muted)" }}>
        Кастомные ноды `.process-node` + dagre layout ppid→pid. Не рисовать SVG-граф с нуля.
      </p>
    </div>
  );
}
