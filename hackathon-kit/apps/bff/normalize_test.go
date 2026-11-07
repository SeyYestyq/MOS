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
				Pod: &struct {
					Namespace string `json:"namespace"`
					Name      string `json:"name"`
				}{Namespace: "hackathon", Name: "attacker-pod"},
			},
		},
		{
			Time:         "2026-11-07T21:15:01Z",
			Type:         "PROCESS_KPROBE",
			FunctionName: "inet_csk_listen_start",
			Process: &RawProcess{
				ExecID: "cexec1", PID: 42, Binary: "/bin/nc",
				Pod: &struct {
					Namespace string `json:"namespace"`
					Name      string `json:"name"`
				}{Namespace: "hackathon", Name: "attacker-pod"},
			},
			Args: []RawArg{{SockArg: &SockArg{Saddr: "0.0.0.0", Sport: 8888, Family: 2}}},
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
	if len(p.Network) != 1 || p.Network[0].Label != "listen" || p.Network[0].Sport != 8888 {
		t.Fatalf("network: %+v", p.Network)
	}
	found := false
	for _, m := range p.Processes[0].Mitre {
		if m == "TA0011" {
			found = true
		}
	}
	if !found {
		t.Fatalf("mitre missing: %+v", p.Processes[0].Mitre)
	}
}

func TestNormalizeRawFile_JSONArray(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "events.json")
	body := `[
	  {"time":"2026-11-07T21:00:00Z","type":"PROCESS_EXEC",
	   "process":{"exec_id":"a1","pid":1,"binary":"/bin/sh","pod":{"namespace":"hackathon","name":"attacker-pod"}}}
	]`
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	p, err := NormalizeRawFile(path, "hackathon", "attacker-pod", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Processes) != 1 || p.Processes[0].ExecID != "a1" {
		t.Fatalf("%+v", p.Processes)
	}
}
