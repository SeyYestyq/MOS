package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// unwrapRuntimeEvent — один адаптер для REST page item и dump (DATA_MODEL).
// Понимает: плоский RawRREvent, RR RuntimeEvent envelope, сырой Tetragon GetEventsResponse.
func unwrapRuntimeEvent(raw json.RawMessage) (RawRREvent, error) {
	raw = json.RawMessage(strings.TrimSpace(string(raw)))
	if len(raw) == 0 || raw[0] == 'n' { // null
		return RawRREvent{}, fmt.Errorf("empty event")
	}

	// 1) Already flat internal shape
	var flat RawRREvent
	if err := json.Unmarshal(raw, &flat); err == nil && flat.Process != nil && flat.Process.ExecID != "" {
		if flat.Time == "" {
			flat.Time = coerceEventTime(raw)
		}
		flat.Threats = parseThreatsFlexible(raw)
		return flat, nil
	}

	// 2) RR RuntimeEvent: { event: {...}, threats, registered_at? }
	var wrap struct {
		Event         json.RawMessage `json:"event"`
		RegisteredAt  string          `json:"registered_at"`
		Threats       json.RawMessage `json:"threats"`
		SourceEvent   json.RawMessage `json:"source_event"`
	}
	_ = json.Unmarshal(raw, &wrap)

	payload := wrap.Event
	if len(payload) == 0 || string(payload) == "null" {
		// sometimes source_event is a JSON string containing Tetragon event
		if len(wrap.SourceEvent) > 0 && wrap.SourceEvent[0] == '"' {
			var s string
			if json.Unmarshal(wrap.SourceEvent, &s) == nil {
				payload = json.RawMessage(s)
			}
		} else if len(wrap.SourceEvent) > 0 && wrap.SourceEvent[0] == '{' {
			payload = wrap.SourceEvent
		} else {
			payload = raw // maybe top-level Tetragon
		}
	}

	ev, err := unwrapTetragonPayload(payload)
	if err != nil {
		return RawRREvent{}, err
	}
	if ev.Time == "" {
		ev.Time = firstNonEmpty(coerceEventTime(payload), wrap.RegisteredAt)
	}
	if th := parseThreatsFlexible(raw); len(th) > 0 {
		ev.Threats = th
	} else if len(wrap.Threats) > 0 {
		ev.Threats = parseThreatsBytes(wrap.Threats)
	}
	return ev, nil
}

func unwrapTetragonPayload(payload json.RawMessage) (RawRREvent, error) {
	var inner struct {
		Time          string `json:"time"`
		ProcessExec   *struct {
			Process *RawProcess `json:"process"`
			Parent  *RawProcess `json:"parent"`
		} `json:"process_exec"`
		ProcessExit *struct {
			Process *RawProcess `json:"process"`
			Parent  *RawProcess `json:"parent"`
		} `json:"process_exit"`
		ProcessKprobe *struct {
			Process      *RawProcess `json:"process"`
			Parent       *RawProcess `json:"parent"`
			FunctionName string      `json:"function_name"`
			Args         []RawArg    `json:"args"`
		} `json:"process_kprobe"`
	}
	if err := json.Unmarshal(payload, &inner); err != nil {
		return RawRREvent{}, err
	}
	ev := RawRREvent{Time: inner.Time}
	switch {
	case inner.ProcessExec != nil && inner.ProcessExec.Process != nil:
		ev.Type = "PROCESS_EXEC"
		ev.Process = inner.ProcessExec.Process
		ev.Parent = inner.ProcessExec.Parent
	case inner.ProcessExit != nil && inner.ProcessExit.Process != nil:
		ev.Type = "PROCESS_EXIT"
		ev.Process = inner.ProcessExit.Process
		ev.Parent = inner.ProcessExit.Parent
	case inner.ProcessKprobe != nil && inner.ProcessKprobe.Process != nil:
		ev.Type = "PROCESS_KPROBE"
		ev.Process = inner.ProcessKprobe.Process
		ev.Parent = inner.ProcessKprobe.Parent
		ev.FunctionName = inner.ProcessKprobe.FunctionName
		ev.Args = inner.ProcessKprobe.Args
	default:
		return RawRREvent{}, fmt.Errorf("unknown tetragon shape")
	}
	return ev, nil
}

func coerceEventTime(raw json.RawMessage) string {
	var t struct {
		Time string `json:"time"`
		Event struct {
			Time string `json:"time"`
		} `json:"event"`
	}
	_ = json.Unmarshal(raw, &t)
	return firstNonEmpty(t.Event.Time, t.Time)
}

func parseThreatsFlexible(raw json.RawMessage) []RawThreat {
	var wrap struct {
		Threats json.RawMessage `json:"threats"`
	}
	if json.Unmarshal(raw, &wrap) != nil || len(wrap.Threats) == 0 {
		return nil
	}
	return parseThreatsBytes(wrap.Threats)
}

func parseThreatsBytes(b json.RawMessage) []RawThreat {
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	// try strict first
	var strict []RawThreat
	if err := json.Unmarshal(b, &strict); err == nil && len(strict) > 0 {
		return strict
	}
	var objs []map[string]json.RawMessage
	if err := json.Unmarshal(b, &objs); err != nil {
		return nil
	}
	out := make([]RawThreat, 0, len(objs))
	for _, o := range objs {
		th := RawThreat{}
		if v, ok := o["severity"]; ok {
			_ = json.Unmarshal(v, &th.Severity)
		}
		if v, ok := o["reason"]; ok {
			_ = json.Unmarshal(v, &th.Reason)
		}
		if v, ok := o["detector"]; ok {
			th.Detector = flexString(v)
		}
		if v, ok := o["tactics_covered"]; ok {
			th.TacticsCovered = flexTactics(v)
		}
		out = append(out, th)
	}
	return out
}

func flexString(b json.RawMessage) string {
	var s string
	if json.Unmarshal(b, &s) == nil {
		return s
	}
	var obj map[string]any
	if json.Unmarshal(b, &obj) == nil {
		if id, ok := obj["id"].(string); ok {
			return id
		}
		if name, ok := obj["name"].(string); ok {
			return name
		}
	}
	return strings.Trim(string(b), `"`)
}

func flexTactics(b json.RawMessage) []string {
	var ss []string
	if json.Unmarshal(b, &ss) == nil {
		return ss
	}
	var objs []map[string]any
	if json.Unmarshal(b, &objs) != nil {
		return nil
	}
	var out []string
	for _, o := range objs {
		if id, ok := o["id"].(string); ok {
			out = append(out, id)
		} else if t, ok := o["tactic"].(string); ok {
			out = append(out, t)
		}
		if tech, ok := o["techniques"].([]any); ok {
			for _, x := range tech {
				if s, ok := x.(string); ok {
					out = append(out, s)
				} else if m, ok := x.(map[string]any); ok {
					if id, ok := m["id"].(string); ok {
						out = append(out, id)
					}
				}
			}
		}
	}
	return out
}
