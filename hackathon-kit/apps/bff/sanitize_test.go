package main

import (
	"strings"
	"testing"
)

func TestSanitizeArguments_XSS(t *testing.T) {
	in := `nc -l 8888 --comment="<script>alert(1)</script>"`
	out := SanitizeArguments(in)
	if strings.Contains(out, "<script>") || strings.Contains(out, "</script>") {
		t.Fatalf("script tags survived: %q", out)
	}
	if !strings.Contains(out, "nc") || !strings.Contains(out, "8888") {
		t.Fatalf("lost payload content: %q", out)
	}
}

func TestSanitizePayload_Processes(t *testing.T) {
	p := &ActivityPayload{
		Processes: []ProcessNode{{
			ExecID:    "x",
			ID:        "x",
			Arguments: `<img onerror="alert(1)">`,
		}},
	}
	SanitizePayload(p)
	if strings.Contains(p.Processes[0].Arguments, "<img") || strings.Contains(p.Processes[0].Arguments, ">") {
		t.Fatalf("unsafe args: %q", p.Processes[0].Arguments)
	}
}
