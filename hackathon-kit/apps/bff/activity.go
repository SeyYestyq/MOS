package main

// ActivityPayload — ответ /api/activity. Канон: docs/DATA_MODEL.md (эволюция контракта).

type ActivityPayload struct {
	Pod       string          `json:"pod"`
	Namespace string          `json:"namespace"`
	Cluster   string          `json:"cluster,omitempty"`
	Processes []ProcessNode   `json:"processes"`
	Network   []NetEdge       `json:"network"`
	Events    []ActivityEvent `json:"events"`
	Threats   []ThreatHit     `json:"threats"`
}

type ProcessNode struct {
	ID           string   `json:"id"`
	ExecID       string   `json:"execId"`
	ParentExecID *string  `json:"parentExecId"`
	ParentStatus string   `json:"parentStatus,omitempty"` // known | outside_window | unknown
	PID          int      `json:"pid"`
	PPID         int      `json:"ppid,omitempty"`
	Binary       string   `json:"binary"`
	Arguments    string   `json:"arguments,omitempty"`
	Cwd          string   `json:"cwd,omitempty"`
	Role         string   `json:"role,omitempty"`
	InInitTree   bool     `json:"inInitTree,omitempty"`
	StartTime    string   `json:"startTime,omitempty"`    // startedAt alias for UI
	StartedAt    string   `json:"startedAt,omitempty"`
	ExitedAt     string   `json:"exitedAt,omitempty"`
	FirstSeen    string   `json:"firstSeen,omitempty"` // firstObservedAt
	LastSeen     string   `json:"lastSeen,omitempty"`  // lastObservedAt — NOT exit
	Mitre        []string `json:"mitre,omitempty"`
	Severity     string   `json:"severity,omitempty"`
	Reason       string   `json:"reason,omitempty"`
}

type NetEdge struct {
	ID     string `json:"id"`
	ExecID string `json:"execId,omitempty"`
	Kind   string `json:"kind,omitempty"` // connect | listen | close
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

type ThreatHit struct {
	ID         string   `json:"id"`
	DetectorID string   `json:"detectorId,omitempty"`
	Severity   string   `json:"severity,omitempty"`
	Reason     string   `json:"reason,omitempty"`
	Tactics    []string `json:"tactics,omitempty"`
	Techniques []string `json:"techniques,omitempty"`
	EventID    string   `json:"eventId,omitempty"`
	ExecID     string   `json:"execId,omitempty"`
	ObservedAt string   `json:"observedAt,omitempty"`
}

func ensureSlices(p *ActivityPayload) {
	if p.Processes == nil {
		p.Processes = []ProcessNode{}
	}
	if p.Network == nil {
		p.Network = []NetEdge{}
	}
	if p.Events == nil {
		p.Events = []ActivityEvent{}
	}
	if p.Threats == nil {
		p.Threats = []ThreatHit{}
	}
}
