package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNormalizeEvents_ExecIDAndSock(t *testing.T) {
	parent := "pexec1"
	events := []RawRREvent{
		{
			Time: "2026-11-07T21:15:00Z",
			Type: "PROCESS_EXEC",
			Process: &RawProcess{
				ExecID: "cexec1", ParentExecID: parent, PID: 42,
				Binary: "/bin/nc", Arguments: "-l -p 8888",
				Pod: &RawPod{Namespace: "hackathon", Name: "attacker-pod"},
			},
		},
		{
			Time:         "2026-11-07T21:15:01Z",
			Type:         "PROCESS_KPROBE",
			FunctionName: "inet_csk_listen_start",
			Process: &RawProcess{
				ExecID: "cexec1", PID: 42, Binary: "/bin/nc",
				Pod: &RawPod{Namespace: "hackathon", Name: "attacker-pod"},
			},
			Args: []RawArg{{SockArg: &SockArg{Saddr: "0.0.0.0", Sport: 8888, Family: "AF_INET"}}},
			Threats: []RawThreat{{
				Severity: "CRITICAL", Reason: "Listening on port 8888",
				TacticsCovered: []string{"TA0011", "T1571"},
			}},
		},
	}
	p, err := NormalizeEvents(events, "hackathon", "attacker-pod", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Processes) != 1 {
		t.Fatalf("procs=%d", len(p.Processes))
	}
	if p.Processes[0].ParentExecID == nil || *p.Processes[0].ParentExecID != parent {
		t.Fatalf("parent edge broken: %+v", p.Processes[0])
	}
	if p.Processes[0].ParentStatus != "outside_window" {
		t.Fatalf("expected outside_window, got %s", p.Processes[0].ParentStatus)
	}
	if len(p.Network) != 1 || p.Network[0].Kind != "listen" || p.Network[0].Sport != 8888 {
		t.Fatalf("network: %+v", p.Network)
	}
	if len(p.Threats) != 1 {
		t.Fatalf("threats=%d", len(p.Threats))
	}
}

func TestNormalizeRawFile_RREnvelope(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "events.json")
	body := `[
	  {
	    "threats": [{"severity":"CRITICAL","reason":"Listening on port 8888","detector":{"id":"TEST"},"tactics_covered":[{"id":"TA0011"}]}],
	    "event": {
	      "time": "2026-11-07T21:15:01.000Z",
	      "process_kprobe": {
	        "function_name": "inet_csk_listen_start",
	        "process": {
	          "exec_id": "e1", "pid": 7, "binary": "/bin/nc",
	          "pod": {"namespace":"hackathon","name":"attacker-pod"}
	        },
	        "args": [{"sock_arg": {"family":"AF_INET","saddr":"0.0.0.0","sport":8888}}]
	      }
	    }
	  }
	]`
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	p, err := NormalizeRawFile(path, "hackathon", "attacker-pod", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Processes) != 1 || p.Processes[0].ExecID != "e1" {
		t.Fatalf("procs=%+v", p.Processes)
	}
	if p.Events[0].T == "" {
		t.Fatal("lost event time")
	}
	if len(p.Network) != 1 || p.Network[0].Kind != "listen" {
		t.Fatalf("net=%+v", p.Network)
	}
}

func TestNormalize_MinMaxTimes(t *testing.T) {
	events := []RawRREvent{
		{Time: "2026-11-07T21:30:00Z", Type: "PROCESS_EXIT", Process: &RawProcess{
			ExecID: "x", PID: 1, Binary: "/bin/sh", Pod: &RawPod{Namespace: "hackathon", Name: "attacker-pod"},
		}},
		{Time: "2026-11-07T21:00:00Z", Type: "PROCESS_EXEC", Process: &RawProcess{
			ExecID: "x", PID: 1, Binary: "/bin/sh", StartTime: "2026-11-07T21:00:00Z",
			Pod: &RawPod{Namespace: "hackathon", Name: "attacker-pod"},
		}},
	}
	p, err := NormalizeEvents(events, "hackathon", "attacker-pod", nil)
	if err != nil {
		t.Fatal(err)
	}
	n := p.Processes[0]
	if n.FirstSeen != "2026-11-07T21:00:00Z" || n.LastSeen != "2026-11-07T21:30:00Z" {
		t.Fatalf("first/last wrong: %s %s", n.FirstSeen, n.LastSeen)
	}
	if n.ExitedAt != "2026-11-07T21:30:00Z" {
		t.Fatalf("exitedAt=%s", n.ExitedAt)
	}
}

func TestNormalize_EmptyArraysNotNull(t *testing.T) {
	p, err := NormalizeEvents(nil, "hackathon", "attacker-pod", nil)
	if err != nil {
		t.Fatal(err)
	}
	if p.Processes == nil || p.Network == nil || p.Events == nil || p.Threats == nil {
		t.Fatalf("nil slices: %+v", p)
	}
}
