package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// RawRREvent is a minimal Tetragon/RR event shape for dump → ActivityPayload.
// Denis extends fields as real dumps arrive; Damir feeds fixtures/raw/*.json.
type RawRREvent struct {
	Time    string          `json:"time"`
	Type    string          `json:"type"` // PROCESS_EXEC | PROCESS_EXIT | PROCESS_KPROBE | ...
	Process *RawProcess     `json:"process"`
	Parent  *RawProcess     `json:"parent"`
	FunctionName string     `json:"function_name"`
	Args    []RawArg        `json:"args"`
	Threats []RawThreat     `json:"threats"`
	Raw     json.RawMessage `json:"-"`
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
	InInitTree   bool   `json:"in_init_tree"`
	Pod          *struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	} `json:"pod"`
}

type RawArg struct {
	SockArg *SockArg `json:"sock_arg"`
	FileArg *struct {
		Path string `json:"path"`
	} `json:"file_arg"`
}

type SockArg struct {
	Family   int    `json:"family"`
	Type     int    `json:"type"`
	Protocol int    `json:"protocol"`
	Saddr    string `json:"saddr"`
	Daddr    string `json:"daddr"`
	Sport    uint32 `json:"sport"`
	Dport    uint32 `json:"dport"`
}

type RawThreat struct {
	Severity       string   `json:"severity"`
	Reason         string   `json:"reason"`
	TacticsCovered []string `json:"tactics_covered"`
	Detector       string   `json:"detector"`
}

// NormalizeRawFile reads a JSON array or NDJSON of RawRREvent and builds ActivityPayload.
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

func parseRawEvents(b []byte) ([]RawRREvent, error) {
	b = bytesTrimSpace(b)
	if len(b) == 0 {
		return nil, fmt.Errorf("empty raw dump")
	}
	if b[0] == '[' {
		var arr []RawRREvent
		if err := json.Unmarshal(b, &arr); err != nil {
			return nil, err
		}
		return arr, nil
	}
	// NDJSON
	var out []RawRREvent
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var ev RawRREvent
		if err := json.Unmarshal([]byte(line), &ev); err != nil {
			return nil, fmt.Errorf("ndjson: %w", err)
		}
		out = append(out, ev)
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
		Processes: nil,
		Network:   nil,
		Events:    nil,
	}
	procs := map[string]*ProcessNode{}

	for i, ev := range events {
		if notBefore != nil && ev.Time != "" {
			if t, err := time.Parse(time.RFC3339Nano, ev.Time); err == nil && t.Before(*notBefore) {
				continue
			}
		}
		if ev.Process == nil || ev.Process.ExecID == "" {
			continue
		}
		if namespace != "" && ev.Process.Pod != nil && ev.Process.Pod.Namespace != "" && ev.Process.Pod.Namespace != namespace {
			continue
		}
		if pod != "" && ev.Process.Pod != nil && ev.Process.Pod.Name != "" && ev.Process.Pod.Name != pod {
			continue
		}

		execID := ev.Process.ExecID
		node, ok := procs[execID]
		if !ok {
			var parent *string
			pe := ev.Process.ParentExecID
			if pe == "" && ev.Parent != nil {
				pe = ev.Parent.ExecID
			}
			if pe != "" {
				parent = &pe
			}
			node = &ProcessNode{
				ID:           execID,
				ExecID:       execID,
				ParentExecID: parent,
				PID:          int(ev.Process.PID),
				Binary:       ev.Process.Binary,
				Arguments:    ev.Process.Arguments,
				Cwd:          ev.Process.Cwd,
				InInitTree:   ev.Process.InInitTree,
				StartTime:    firstNonEmpty(ev.Process.StartTime, ev.Time),
				FirstSeen:    ev.Time,
				LastSeen:     ev.Time,
				Mitre:        nil,
			}
			procs[execID] = node
		} else {
			if ev.Time != "" {
				node.LastSeen = ev.Time
			}
		}

		for _, th := range ev.Threats {
			if th.Severity != "" {
				node.Severity = th.Severity
			}
			if th.Reason != "" {
				node.Reason = th.Reason
			}
			node.Mitre = appendUnique(node.Mitre, th.TacticsCovered...)
			if looksLikeNC(node.Binary) && (containsPort(th.Reason, "8888") || ev.FunctionName == "inet_csk_listen_start") {
				node.Mitre = appendUnique(node.Mitre, "TA0011", "T1571")
				node.Role = "suspect"
			}
		}

		summary := ev.Type
		if ev.Process.Binary != "" {
			summary = ev.Type + " " + ev.Process.Binary
		}
		reason := ""
		if len(ev.Threats) > 0 {
			reason = ev.Threats[0].Reason
		}
		p.Events = append(p.Events, ActivityEvent{
			ID:      fmt.Sprintf("ev-%d-%s", i, execID),
			T:       ev.Time,
			Type:    ev.Type,
			ExecID:  execID,
			PID:     int(ev.Process.PID),
			Summary: summary,
			Reason:  reason,
		})

		if strings.Contains(strings.ToUpper(ev.Type), "KPROBE") || ev.FunctionName != "" {
			for j, arg := range ev.Args {
				if arg.SockArg == nil {
					continue
				}
				sa := arg.SockArg
				label := "egress"
				if ev.FunctionName == "inet_csk_listen_start" || sa.Dport == 0 && sa.Sport > 0 {
					label = "listen"
				}
				p.Network = append(p.Network, NetEdge{
					ID:     fmt.Sprintf("net-%d-%d", i, j),
					ExecID: execID,
					Saddr:  sa.Saddr,
					Sport:  int(sa.Sport),
					Daddr:  sa.Daddr,
					Dport:  int(sa.Dport),
					Family: "AF_INET",
					Dst:    sa.Daddr,
					Port:   int(sa.Dport),
					Proto:  "tcp",
					Label:  label,
					T:      ev.Time,
				})
			}
		}
	}

	for _, n := range procs {
		p.Processes = append(p.Processes, *n)
	}
	SanitizePayload(p)
	if err := validateDAG(p); err != nil {
		return nil, err
	}
	return p, nil
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
