package main

import "testing"

func TestLoadFixture_Sample(t *testing.T) {
	p, err := loadFixture("fixtures/sample-activity.json")
	if err != nil {
		t.Fatal(err)
	}
	if p.Pod == "" || len(p.Processes) == 0 {
		t.Fatalf("empty fixture: %+v", p)
	}
	byID := map[string]ProcessNode{}
	for _, n := range p.Processes {
		if n.ExecID == "" {
			t.Fatalf("missing execId on %#v", n)
		}
		byID[n.ExecID] = n
	}
	// edges only via parentExecId
	for _, n := range p.Processes {
		if n.ParentExecID == nil || *n.ParentExecID == "" {
			continue
		}
		if _, ok := byID[*n.ParentExecID]; !ok {
			t.Fatalf("orphan parentExecId %s for %s", *n.ParentExecID, n.ExecID)
		}
	}
	if len(p.Network) == 0 {
		t.Fatal("expected network edges in sample")
	}
	if len(p.Events) == 0 {
		t.Fatal("expected events in sample")
	}
}

func TestValidateDAG_RejectsMissingExecID(t *testing.T) {
	err := validateDAG(&ActivityPayload{
		Processes: []ProcessNode{{ID: "a", PID: 1}},
	})
	if err == nil {
		t.Fatal("expected error")
	}
}
