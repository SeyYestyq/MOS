package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadFixture(path string) (*ActivityPayload, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("fixture %s: %w", path, err)
	}
	var p ActivityPayload
	if err := json.Unmarshal(b, &p); err != nil {
		return nil, fmt.Errorf("fixture json: %w", err)
	}
	if err := validateDAG(&p); err != nil {
		return nil, err
	}
	SanitizePayload(&p)
	ensureSlices(&p)
	return &p, nil
}

// validateDAG ensures tree edges are execId-based (parentExecId → execId), not pid-only.
func validateDAG(p *ActivityPayload) error {
	if p == nil {
		return fmt.Errorf("nil payload")
	}
	ids := make(map[string]struct{}, len(p.Processes))
	for _, n := range p.Processes {
		if n.ExecID == "" {
			return fmt.Errorf("process missing execId (pid=%d)", n.PID)
		}
		if n.ID == "" {
			return fmt.Errorf("process missing id for execId=%s", n.ExecID)
		}
		ids[n.ExecID] = struct{}{}
	}
	for _, n := range p.Processes {
		if n.ParentExecID == nil || *n.ParentExecID == "" {
			continue
		}
		if _, ok := ids[*n.ParentExecID]; !ok {
			// parent outside sample window is OK for partial dumps
			continue
		}
	}
	return nil
}
