/** Контракт UI ↔ BFF. Дерево связывать по exec_id, не только pid. */

export type ProcessNode = {
  id: string; // обычно = exec_id
  execId: string;
  parentExecId?: string | null;
  pid: number;
  ppid?: number;
  binary: string;
  arguments?: string;
  cwd?: string;
  role?: "entry" | "shell" | "tool" | "suspect" | string;
  inInitTree?: boolean;
  startTime?: string;
  firstSeen?: string;
  lastSeen?: string;
  mitre?: string[]; // e.g. TA0011, T1571
};

export type NetEdge = {
  id: string;
  execId?: string;
  fromPid?: number;
  saddr?: string;
  sport?: number;
  daddr?: string;
  dport?: number;
  family?: string;
  dst?: string; // legacy shorthand host
  port?: number;
  proto?: string;
  label?: string;
  t?: string;
};

export type ActivityEvent = {
  id: string;
  t: string;
  type: string; // PROCESS_EXEC | PROCESS_EXIT | PROCESS_KPROBE | ...
  execId?: string;
  pid?: number;
  summary: string;
  reason?: string;
  raw?: unknown;
};

export type ActivityPayload = {
  pod: string;
  namespace: string;
  cluster?: string;
  processes: ProcessNode[];
  network: NetEdge[];
  events: ActivityEvent[];
};
