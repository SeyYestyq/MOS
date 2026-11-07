package main

import (
	"strings"
	"unicode"
)

// SanitizeArguments strips HTML/script-ish payloads from process cmdline (XSS guard).
func SanitizeArguments(s string) string {
	if s == "" {
		return s
	}
	replacer := strings.NewReplacer(
		"<", "‹",
		">", "›",
		`"`, "'",
		"`", "'",
	)
	out := replacer.Replace(s)
	var b strings.Builder
	b.Grow(len(out))
	for _, r := range out {
		if r == '\u0000' || (unicode.IsControl(r) && r != '\t' && r != '\n' && r != '\r') {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func SanitizePayload(p *ActivityPayload) {
	if p == nil {
		return
	}
	for i := range p.Processes {
		p.Processes[i].Arguments = SanitizeArguments(p.Processes[i].Arguments)
		p.Processes[i].Binary = SanitizeArguments(p.Processes[i].Binary)
	}
	for i := range p.Events {
		p.Events[i].Summary = SanitizeArguments(p.Events[i].Summary)
		p.Events[i].Reason = SanitizeArguments(p.Events[i].Reason)
	}
}
