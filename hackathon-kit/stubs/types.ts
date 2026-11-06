export type ProcessNode = {
  id: string;
  pid: number;
  ppid: number;
  binary: string;
  role?: "entry" | "shell" | "tool" | "suspect" | string;
  cmdline?: string;
  firstSeen?: string;
  lastSeen?: string;
};

export type NetEdge = {
  id: string;
  fromPid?: number;
  dst: string;
  port?: number;
  proto?: string;
  label?: string;
  t?: string;
};

export type ActivityEvent = {
  id: string;
  t: string;
  type: string;
  pid?: number;
  summary: string;
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
