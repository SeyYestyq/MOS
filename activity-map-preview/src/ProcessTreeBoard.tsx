import { useMemo } from "react";
import type { ProcessNode } from "./mockData";

const NODE_W = 178;
const NODE_H = 68;
const COL_GAP = 210;
const ROW_GAP = 92;
const PAD_X = 48;
const PAD_Y = 56;

const ROLE_COLOR: Record<ProcessNode["role"], string> = {
  entry: "#37d5ff",
  shell: "#ffce63",
  tool: "#84cc16",
  suspect: "#ff5f8d",
};

const ROLE_LABEL: Record<ProcessNode["role"], string> = {
  entry: "entry",
  shell: "shell",
  tool: "tool",
  suspect: "suspect",
};

type Pos = { x: number; y: number };

function shortName(binary: string): string {
  const base = binary.split("/").pop() || binary;
  return base.length > 16 ? `${base.slice(0, 15)}…` : base;
}

function layoutTree(nodes: ProcessNode[]) {
  const byPid = new Map(nodes.map((n) => [n.pid, n]));
  const children = new Map<number, ProcessNode[]>();
  let root: ProcessNode | undefined;

  for (const n of nodes) {
    if (!byPid.has(n.ppid) || n.ppid === 0) {
      root = root ?? n;
      continue;
    }
    const list = children.get(n.ppid) || [];
    list.push(n);
    children.set(n.ppid, list);
  }
  if (!root) root = nodes[0];

  const positions: Record<string, Pos> = {};
  const depths: Record<string, number> = {};
  let leaf = 0;

  function walk(node: ProcessNode, depth: number) {
    const kids = children.get(node.pid) || [];
    depths[node.id] = depth;
    if (!kids.length) {
      positions[node.id] = {
        x: PAD_X + depth * COL_GAP,
        y: PAD_Y + leaf * ROW_GAP,
      };
      leaf += 1;
      return;
    }
    for (const kid of kids) walk(kid, depth + 1);
    const first = positions[kids[0].id];
    const last = positions[kids[kids.length - 1].id];
    positions[node.id] = {
      x: PAD_X + depth * COL_GAP,
      y: (first.y + last.y) / 2,
    };
  }

  walk(root, 0);

  // orphans not under root
  for (const n of nodes) {
    if (!positions[n.id]) {
      positions[n.id] = {
        x: PAD_X,
        y: PAD_Y + leaf * ROW_GAP,
      };
      depths[n.id] = 0;
      leaf += 1;
    }
  }

  const maxDepth = Math.max(0, ...Object.values(depths));
  const svgWidth = Math.max(720, PAD_X * 2 + (maxDepth + 1) * COL_GAP);
  const svgHeight = Math.max(360, PAD_Y * 2 + leaf * ROW_GAP);

  const edges: Array<{ from: string; to: string }> = [];
  for (const n of nodes) {
    const parent = byPid.get(n.ppid);
    if (parent) edges.push({ from: parent.id, to: n.id });
  }

  return { positions, edges, svgWidth, svgHeight, maxDepth };
}

function edgePath(source: Pos, target: Pos): string {
  const sx = source.x + NODE_W;
  const sy = source.y + NODE_H / 2;
  const tx = target.x;
  const ty = target.y + NODE_H / 2;
  const mid = (sx + tx) / 2;
  return `M ${sx} ${sy} C ${mid} ${sy}, ${mid} ${ty}, ${tx} ${ty}`;
}

type Props = {
  processes: ProcessNode[];
  selectedId: string | null;
  onSelect: (node: ProcessNode) => void;
};

export function ProcessTreeBoard({ processes, selectedId, onSelect }: Props) {
  const { positions, edges, svgWidth, svgHeight, maxDepth } = useMemo(
    () => layoutTree(processes),
    [processes],
  );

  const byId = useMemo(
    () => Object.fromEntries(processes.map((p) => [p.id, p])),
    [processes],
  );

  return (
    <div className="admin-topology-canvas">
      <svg
        className="admin-topology"
        width={svgWidth}
        height={svgHeight}
        viewBox={`0 0 ${svgWidth} ${svgHeight}`}
        role="img"
        aria-label="Дерево процессов контейнера"
      >
        <defs>
          <marker
            id="tree-arrow"
            markerWidth="7"
            markerHeight="5"
            refX="7"
            refY="2.5"
            orient="auto"
          >
            <path d="M 0 0 L 7 2.5 L 0 5 Z" fill="#6b7280" fillOpacity="0.85" />
          </marker>
        </defs>

        {Array.from({ length: maxDepth + 1 }, (_, depth) => (
          <text
            key={depth}
            x={PAD_X + depth * COL_GAP}
            y={28}
            className="admin-topology-layer"
          >
            {depth === 0 ? "ROOT" : `L${depth}`}
          </text>
        ))}

        {edges.map((edge) => {
          const source = positions[edge.from];
          const target = positions[edge.to];
          if (!source || !target) return null;
          const child = byId[edge.to];
          const color = child ? ROLE_COLOR[child.role] : "#6b7280";
          return (
            <path
              key={`${edge.from}-${edge.to}`}
              d={edgePath(source, target)}
              fill="none"
              stroke={color}
              strokeWidth="1.8"
              strokeOpacity={0.55}
              markerEnd="url(#tree-arrow)"
              className="admin-topology-edge-path"
            />
          );
        })}

        {processes.map((node) => {
          const pos = positions[node.id];
          if (!pos) return null;
          const color = ROLE_COLOR[node.role];
          const selected = selectedId === node.id;
          return (
            <g
              key={node.id}
              className={`admin-topology-node ${selected ? "selected" : ""}`}
              onClick={() => onSelect(node)}
              style={{ cursor: "pointer" }}
            >
              {selected ? (
                <rect
                  x={pos.x - 4}
                  y={pos.y - 4}
                  width={NODE_W + 8}
                  height={NODE_H + 8}
                  rx={14}
                  fill="none"
                  stroke={color}
                  strokeWidth="2"
                  className="admin-topology-node-ring"
                />
              ) : null}
              <rect
                x={pos.x}
                y={pos.y}
                width={NODE_W}
                height={NODE_H}
                rx={12}
                className="admin-topology-node-card"
                stroke={color}
              />
              <circle cx={pos.x + 15} cy={pos.y + 24} r="4" fill={color} />
              <text x={pos.x + 27} y={pos.y + 28} className="admin-topology-name">
                {shortName(node.binary)}
              </text>
              <text x={pos.x + 15} y={pos.y + 48} className="admin-topology-port">
                pid {node.pid}
                {node.ppid ? ` ← ${node.ppid}` : ""}
              </text>
              <text
                x={pos.x + NODE_W - 10}
                y={pos.y + 48}
                textAnchor="end"
                className="admin-topology-status"
                fill={color}
              >
                {ROLE_LABEL[node.role]}
              </text>
            </g>
          );
        })}
      </svg>

      <div className="admin-topology-legend">
        <span className="admin-topology-legend-title">Роли:</span>
        {(Object.keys(ROLE_COLOR) as Array<ProcessNode["role"]>).map((role) => (
          <span key={role} className="admin-topology-legend-item">
            <span
              className="admin-topology-legend-dot"
              style={{ background: ROLE_COLOR[role] }}
            />
            {ROLE_LABEL[role]}
          </span>
        ))}
      </div>
    </div>
  );
}
