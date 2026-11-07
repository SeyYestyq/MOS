package main

// ActivityPayload mirrors hackathon-kit/stubs/types.ts — keep in sync with Angelina.

type ActivityPayload struct {
	Pod       string         `json:"pod"`
	Namespace string         `json:"namespace"`
	Cluster   string         `json:"cluster,omitempty"`
	Processes []ProcessNode  `json:"processes"`
	Network   []NetEdge      `json:"network"`
	Events    []ActivityEvent `json:"events"`
}

type ProcessNode struct {
	ID           string   `json:"id"`
	ExecID       string   `json:"execId"`
	ParentExecID *string  `json:"parentExecId"`
	PID          int      `json:"pid"`
	PPID         int      `json:"ppid,omitempty"`
	Binary       string   `json:"binary"`
	Arguments    string   `json:"arguments,omitempty"`
	Cwd          string   `json:"cwd,omitempty"`
	Role         string   `json:"role,omitempty"`
	InInitTree   bool     `json:"inInitTree,omitempty"`
	StartTime    string   `json:"startTime,omitempty"`
	FirstSeen    string   `json:"firstSeen,omitempty"`
	LastSeen     string   `json:"lastSeen,omitempty"`
	Mitre        []string `json:"mitre,omitempty"`
	Severity     string   `json:"severity,omitempty"`
	Reason       string   `json:"reason,omitempty"`
}

type NetEdge struct {
	ID     string `json:"id"`
	ExecID string `json:"execId,omitempty"`
	FromPID int   `json:"fromPid,omitempty"`
	Saddr  string `json:"saddr,omitempty"`
	Sport  int    `json:"sport,omitempty"`
	Daddr  string `json:"daddr,omitempty"`
	Dport  int    `json:"dport,omitempty"`
	Family string `json:"family,omitempty"`
	Dst    string `json:"dst,omitempty"`
	Port   int    `json:"port,omitempty"`
	Proto  string `json:"proto,omitempty"`
	Label  string `json:"label,omitempty"`
	T      string `json:"t,omitempty"`
}

type ActivityEvent struct {
	ID      string `json:"id"`
	T       string `json:"t"`
	Type    string `json:"type"`
	ExecID  string `json:"execId,omitempty"`
	PID     int    `json:"pid,omitempty"`
	Summary string `json:"summary"`
	Reason  string `json:"reason,omitempty"`
}
