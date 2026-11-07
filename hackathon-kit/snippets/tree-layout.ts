/**
 * Алгоритм layout дерева процессов (для использования в web).
 * Поддерживает strict DAG по execId / parentExecId, а также fallback по pid / ppid.
 */

export type TreeNode = {
  id: string;
  execId?: string;
  parentExecId?: string | null;
  pid?: number;
  ppid?: number;
};

export type Pos = { x: number; y: number };

const COL_GAP = 220;
const ROW_GAP = 96;
const PAD_X = 48;
const PAD_Y = 56;

export function layoutProcessTree(nodes: TreeNode[]) {
  // Идентифицируем ноды по execId (если есть) или по pid/id
  const nodeKey = (n: TreeNode) => n.execId || String(n.pid ?? n.id);
  const parentKey = (n: TreeNode) => n.parentExecId || (n.ppid ? String(n.ppid) : null);

  const byKey = new Map<string, TreeNode>();
  for (const n of nodes) {
    byKey.set(nodeKey(n), n);
  }

  const children = new Map<string, TreeNode[]>();
  let root: TreeNode | undefined;

  for (const n of nodes) {
    const pKey = parentKey(n);
    if (!pKey || !byKey.has(pKey) || pKey === "0") {
      root = root ?? n;
      continue;
    }
    const list = children.get(pKey) || [];
    list.push(n);
    children.set(pKey, list);
  }

  if (!root && nodes.length > 0) {
    root = nodes[0];
  }

  const positions: Record<string, Pos> = {};
  let leaf = 0;

  function walk(node: TreeNode, depth: number) {
    const k = nodeKey(node);
    const kids = children.get(k) || [];
    if (!kids.length) {
      positions[node.id] = { x: PAD_X + depth * COL_GAP, y: PAD_Y + leaf * ROW_GAP };
      leaf += 1;
      return;
    }
    for (const kid of kids) {
      walk(kid, depth + 1);
    }
    const first = positions[kids[0].id];
    const last = positions[kids[kids.length - 1].id];
    positions[node.id] = {
      x: PAD_X + depth * COL_GAP,
      y: (first.y + last.y) / 2,
    };
  }

  if (root) walk(root, 0);

  // Для изолированных нод (если дерево распалось)
  for (const n of nodes) {
    if (!positions[n.id]) {
      positions[n.id] = { x: PAD_X, y: PAD_Y + leaf * ROW_GAP };
      leaf += 1;
    }
  }

  const edges = nodes
    .filter((n) => {
      const pKey = parentKey(n);
      return pKey && byKey.has(pKey);
    })
    .map((n) => ({
      from: byKey.get(parentKey(n)!)!.id,
      to: n.id,
    }));

  return { positions, edges, leafCount: leaf };
}

export function edgePath(source: Pos, target: Pos, nodeW = 190, nodeH = 72): string {
  const sx = source.x + nodeW;
  const sy = source.y + nodeH / 2;
  const tx = target.x;
  const ty = target.y + nodeH / 2;
  const mid = (sx + tx) / 2;
  return `M ${sx} ${sy} C ${mid} ${sy}, ${mid} ${ty}, ${tx} ${ty}`;
}
