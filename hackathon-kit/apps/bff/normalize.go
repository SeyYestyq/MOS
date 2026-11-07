package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type RawRREvent struct {
	Time         string      `json:"time"`
	Type         string      `json:"type"`
	Process      *RawProcess `json:"process"`
	Parent       *RawProcess `json:"parent"`
	FunctionName string      `json:"function_name"`
	Args         []RawArg    `json:"args"`
	Threats      []RawThreat `json:"threats"`
}

type RawProcess struct {
	ExecID       string `json:"exec_id"`
	ParentExecID string `json:"parent_exec_id"`
	PID          uint32 `json:"pid"`
	UID          uint32 `json:"uid"`
	Binary       string `json:"binary"`
	Arguments    string `json:"arguments"`
	Cwd          string `json:"cwd"`
	StartTime    string `json:"start_time"`
	InInitTree   bool    `json:"in_init_tree"`
	Pod          *RawPod `json:"pod"`
}

type RawPod struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	UID       string `json:"uid"`
	Container *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"container"`
}

type RawArg struct {
	SockArg *SockArg `json:"sock_arg"`
	FileArg *struct {
		Path string `json:"path"`
	} `json:"file_arg"`
}

// SockArg — Tetragon KprobeSock (family/protocol обычно строки).
type SockArg struct {
	Family   flexStr `json:"family"`
	Type     flexStr `json:"type"`
	Protocol flexStr `json:"protocol"`
	Saddr    string  `json:"saddr"`
	Daddr    string  `json:"daddr"`
	Sport    uint32  `json:"sport"`
	Dport    uint32  `json:"dport"`
}

type flexStr string

func (f *flexStr) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		*f = ""
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*f = flexStr(s)
		return nil
	}
	*f = flexStr(string(b))
	return nil
}

type RawThreat struct {
	Severity       string   `json:"severity"`
	Reason         string   `json:"reason"`
	TacticsCovered []string `json:"tactics_covered"`
	Detector       string   `json:"detector"`
}

func NormalizeRawFile(path, namespace, pod string, notBefore *time.Time) (*ActivityPayload, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	events, err := parseRawEvents(b)
	if err != nil {
		return nil, err
	}
	return NormalizeEvents(events, namespace, pod, notBefore)
}

// parseRawEvents — dump/API эквивалентны: каждый элемент через unwrapRuntimeEvent.
func parseRawEvents(b []byte) ([]RawRREvent, error) {
	b = bytesTrimSpace(b)
	if len(b) == 0 {
		return nil, fmt.Errorf("empty raw dump")
	}
	var items []json.RawMessage
	if b[0] == '[' {
		if err := json.Unmarshal(b, &items); err != nil {
			return nil, err
		}
	} else {
		for _, line := range strings.Split(string(b), "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			items = append(items, json.RawMessage(line))
		}
	}
	out := make([]RawRREvent, 0, len(items))
	var skipped int
	for _, it := range items {
		ev, err := unwrapRuntimeEvent(it)
		if err != nil {
			skipped++
			continue
		}
		out = append(out, ev)
	}
	if len(out) == 0 && len(items) > 0 {
		return nil, fmt.Errorf("raw dump: 0 events decoded (%d skipped)", skipped)
	}
	return out, nil
}

func bytesTrimSpace(b []byte) []byte {
	return []byte(strings.TrimSpace(string(b)))
}

func NormalizeEvents(events []RawRREvent, namespace, pod string, notBefore *time.Time) (*ActivityPayload, error) {
	p := &ActivityPayload{
		Pod:       pod,
		Namespace: namespace,
		Processes: []ProcessNode{},
		Network:   []NetEdge{},
		Events:    []ActivityEvent{},
		Threats:   []ThreatHit{},
	}
	procs := map[string]*ProcessNode{}

	for i, ev := range events {
		if notBefore != nil && ev.Time != "" {
			if t, err := parseTS(ev.Time); err == nil && t.Before(*notBefore) {
				continue
			}
		}
		if ev.Process == nil || ev.Process.ExecID == "" {
			continue
		}
		// Strict pod filter: unknown pod does NOT match a requested pod
		if pod != "" || namespace != "" {
			if ev.Process.Pod == nil {
				continue
			}
			if namespace != "" && ev.Process.Pod.Namespace != namespace {
				continue
			}
			if pod != "" && ev.Process.Pod.Name != pod {
				continue
			}
		}

		execID := ev.Process.ExecID
		node, ok := procs[execID]
		if !ok {
			var parent *string
			pe := ev.Process.ParentExecID
			if pe == "" && ev.Parent != nil {
				pe = ev.Parent.ExecID
			}
			status := "unknown"
			if pe != "" {
				parent = &pe
				status = "outside_window" // until we know parent is in snapshot
			}
			started := firstNonEmpty(ev.Process.StartTime, "")
			node = &ProcessNode{
				ID:           execID,
				ExecID:       execID,
				ParentExecID: parent,
				ParentStatus: status,
				PID:          int(ev.Process.PID),
				Binary:       ev.Process.Binary,
				Arguments:    ev.Process.Arguments,
				Cwd:          ev.Process.Cwd,
				InInitTree:   ev.Process.InInitTree,
				StartTime:    started,
				StartedAt:    started,
				FirstSeen:    ev.Time,
				LastSeen:     ev.Time,
				Mitre:        []string{},
			}
			procs[execID] = node
		} else {
			// enrich incomplete first sighting
			if node.Binary == "" && ev.Process.Binary != "" {
				node.Binary = ev.Process.Binary
			}
			if node.Arguments == "" && ev.Process.Arguments != "" {
				node.Arguments = ev.Process.Arguments
			}
			if node.Cwd == "" && ev.Process.Cwd != "" {
				node.Cwd = ev.Process.Cwd
			}
			if node.StartedAt == "" && ev.Process.StartTime != "" {
				node.StartedAt = ev.Process.StartTime
				node.StartTime = ev.Process.StartTime
			}
			if node.ParentExecID == nil {
				pe := ev.Process.ParentExecID
				if pe == "" && ev.Parent != nil {
					pe = ev.Parent.ExecID
				}
				if pe != "" {
					node.ParentExecID = &pe
					node.ParentStatus = "outside_window"
				}
			}
			node.FirstSeen = minTS(node.FirstSeen, ev.Time)
			node.LastSeen = maxTS(node.LastSeen, ev.Time)
		}

		if strings.EqualFold(ev.Type, "PROCESS_EXIT") && ev.Time != "" {
			node.ExitedAt = ev.Time
		}

		eventID := fmt.Sprintf("ev-%d-%s", i, execID)
		summary := ev.Type
		if ev.Process.Binary != "" {
			summary = ev.Type + " " + ev.Process.Binary
		}
		reason := ""
		if len(ev.Threats) > 0 {
			reason = ev.Threats[0].Reason
		}
		p.Events = append(p.Events, ActivityEvent{
			ID: eventID, T: ev.Time, Type: ev.Type, ExecID: execID,
			PID: int(ev.Process.PID), Summary: summary, Reason: reason,
		})

		for ti, th := range ev.Threats {
			if th.Severity != "" {
				node.Severity = th.Severity
			}
			if th.Reason != "" {
				node.Reason = th.Reason
			}
			node.Mitre = appendUnique(node.Mitre, th.TacticsCovered...)
			p.Threats = append(p.Threats, ThreatHit{
				ID:         fmt.Sprintf("th-%d-%d", i, ti),
				DetectorID: th.Detector,
				Severity:   th.Severity,
				Reason:     th.Reason,
				Tactics:    th.TacticsCovered,
				EventID:    eventID,
				ExecID:     execID,
				ObservedAt: ev.Time,
			})
			if looksLikeNC(node.Binary) && (containsPort(th.Reason, "8888") || ev.FunctionName == "inet_csk_listen_start") {
				node.Mitre = appendUnique(node.Mitre, "TA0011", "T1571")
				node.Role = "suspect"
			}
		}

		if strings.Contains(strings.ToUpper(ev.Type), "KPROBE") || ev.FunctionName != "" {
			for j, arg := range ev.Args {
				if arg.SockArg == nil {
					continue
				}
				sa := arg.SockArg
				kind := netKind(ev.FunctionName, sa)
				fam := string(sa.Family)
				if fam == "" {
					fam = "AF_INET"
				}
				proto := string(sa.Protocol)
				if proto == "" || strings.Contains(strings.ToUpper(proto), "TCP") || proto == "6" {
					proto = "tcp"
				}
				p.Network = append(p.Network, NetEdge{
					ID:     fmt.Sprintf("net-%d-%d", i, j),
					ExecID: execID,
					Kind:   kind,
					Saddr:  sa.Saddr,
					Sport:  int(sa.Sport),
					Daddr:  sa.Daddr,
					Dport:  int(sa.Dport),
					Family: fam,
					Dst:    sa.Daddr,
					Port:   int(sa.Dport),
					Proto:  proto,
					Label:  kind,
					T:      ev.Time,
				})
			}
		}
	}

	// resolve parentStatus
	for _, n := range procs {
		if n.ParentExecID == nil || *n.ParentExecID == "" {
			n.ParentStatus = "unknown"
		} else if _, ok := procs[*n.ParentExecID]; ok {
			n.ParentStatus = "known"
		} else {
			n.ParentStatus = "outside_window"
		}
		p.Processes = append(p.Processes, *n)
	}

	SanitizePayload(p)
	ensureSlices(p)
	if err := validateDAG(p); err != nil {
		return nil, err
	}
	return p, nil
}

func netKind(fn string, sa *SockArg) string {
	fn = strings.ToLower(fn)
	switch {
	case strings.Contains(fn, "listen"):
		return "listen"
	case strings.Contains(fn, "close"):
		return "close"
	case strings.Contains(fn, "connect"):
		return "connect"
	default:
		if sa != nil && sa.Dport == 0 && sa.Sport > 0 {
			return "listen"
		}
		return "connect"
	}
}

func parseTS(s string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, nil
	}
	return time.Parse(time.RFC3339, s)
}

func minTS(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	ta, ea := parseTS(a)
	tb, eb := parseTS(b)
	if ea != nil || eb != nil {
		return a
	}
	if tb.Before(ta) {
		return b
	}
	return a
}

func maxTS(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	ta, ea := parseTS(a)
	tb, eb := parseTS(b)
	if ea != nil || eb != nil {
		return b
	}
	if tb.After(ta) {
		return b
	}
	return a
}

func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func appendUnique(dst []string, add ...string) []string {
	seen := map[string]struct{}{}
	for _, s := range dst {
		seen[s] = struct{}{}
	}
	for _, s := range add {
		if s == "" {
			continue
		}
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		dst = append(dst, s)
	}
	return dst
}

func looksLikeNC(binary string) bool {
	b := strings.ToLower(binary)
	return strings.HasSuffix(b, "/nc") || strings.HasSuffix(b, "/ncat") || strings.Contains(b, "netcat")
}

func containsPort(s, port string) bool {
	return strings.Contains(s, port)
}
