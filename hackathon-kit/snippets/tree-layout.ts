/**
 * Готовый алгоритм layout дерева процессов (скопировать в web завтра).
 * Не UI — только позиции.
 */

export type TreeNode = { id: string; pid: number; ppid: number };

export type Pos = { x: number; y: number };

const COL_GAP = 210;
const ROW_GAP = 92;
const PAD_X = 48;
const PAD_Y = 56;

export function layoutProcessTree(nodes: TreeNode[]) {
  const byPid = new Map(nodes.map((n) => [n.pid, n]));
  const children = new Map<number, TreeNode[]>();
  let root: TreeNode | undefined;

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
  let leaf = 0;

  function walk(node: TreeNode, depth: number) {
    const kids = children.get(node.pid) || [];
    if (!kids.length) {
      positions[node.id] = { x: PAD_X + depth * COL_GAP, y: PAD_Y + leaf * ROW_GAP };
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

  if (root) walk(root, 0);

  for (const n of nodes) {
    if (!positions[n.id]) {
      positions[n.id] = { x: PAD_X, y: PAD_Y + leaf * ROW_GAP };
      leaf += 1;
    }
  }

  const edges = nodes
    .filter((n) => byPid.has(n.ppid))
    .map((n) => ({ from: byPid.get(n.ppid)!.id, to: n.id }));

  return { positions, edges, leafCount: leaf };
}

export function edgePath(source: Pos, target: Pos, nodeW = 178, nodeH = 68): string {
  const sx = source.x + nodeW;
  const sy = source.y + nodeH / 2;
  const tx = target.x;
  const ty = target.y + nodeH / 2;
  const mid = (sx + tx) / 2;
  return `M ${sx} ${sy} C ${mid} ${sy}, ${mid} ${ty}, ${tx} ${ty}`;
}
