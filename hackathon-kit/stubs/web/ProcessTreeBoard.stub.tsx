/**
 * STUB — реализовать завтра.
 * Референс layout/SVG: new4 ServiceTopology.tsx + hackathon-kit/snippets/tree-layout.ts
 * Стили: design/tokens.css + design/topology.css
 */
import type { ProcessNode } from "../types";

export type ProcessTreeBoardProps = {
  processes: ProcessNode[];
  selectedId: string | null;
  onSelect: (node: ProcessNode) => void;
};

export function ProcessTreeBoard(_props: ProcessTreeBoardProps) {
  return (
    <div className="admin-surface liquid-metal-frame">
      <div className="admin-surface-head">
        <h3>Process tree</h3>
        <span>TODO</span>
      </div>
      <p style={{ color: "var(--text-muted)" }}>
        Заготовка: SVG дерево ppid→pid. Портировать рендер узлов из ServiceTopology.
      </p>
    </div>
  );
}
