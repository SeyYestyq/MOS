package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// fetchLiveActivity pulls full events via public-api REST slice (has args/sock_arg).
// ClickHouse path intentionally unused here — no args for network map.
func fetchLiveActivity(namespace, pod string) (*ActivityPayload, error) {
	base := strings.TrimRight(os.Getenv("RR_BASE_URL"), "/")
	if base == "" {
		return nil, fmt.Errorf("RR_BASE_URL empty")
	}
	token := os.Getenv("RR_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("RR_TOKEN empty (set personal Access-Token)")
	}

	sliceSize := 100
	if v := os.Getenv("RR_SLICE_SIZE"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			sliceSize = n
		}
	}
	pages := 5
	if v := os.Getenv("RR_PAGES"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			pages = n
		}
	}

	client := &http.Client{Timeout: 30 * time.Second}
	if os.Getenv("RR_INSECURE_TLS") == "true" || os.Getenv("RR_INSECURE_TLS") == "1" {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // hackathon self-signed
		}
	}

	var all []RawRREvent
	cursor := time.Now().UTC().Format(time.RFC3339Nano)
	for i := 0; i < pages; i++ {
		u, err := url.Parse(base + "/api/v1/public-api/runtime-event/slice/left")
		if err != nil {
			return nil, err
		}
		q := u.Query()
		q.Set("cursor", cursor)
		q.Set("slice_size", strconv.Itoa(sliceSize))
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Access-Token", token)
		req.Header.Set("Accept", "application/json")

		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(res.Body)
		_ = res.Body.Close()
		if err != nil {
			return nil, err
		}
		if res.StatusCode >= 300 {
			return nil, fmt.Errorf("RR slice HTTP %d: %s", res.StatusCode, truncate(string(body), 200))
		}

		var page struct {
			RuntimeEvents []json.RawMessage `json:"runtime_events"`
			LeftCursor    string            `json:"left_cursor"`
		}
		if err := json.Unmarshal(body, &page); err != nil {
			return nil, fmt.Errorf("RR slice decode: %w", err)
		}
		if len(page.RuntimeEvents) == 0 {
			break
		}
		for _, raw := range page.RuntimeEvents {
			ev, err := unwrapRuntimeEvent(raw)
			if err != nil {
				continue
			}
			all = append(all, ev)
		}
		if page.LeftCursor == "" || page.LeftCursor == cursor {
			break
		}
		cursor = page.LeftCursor
	}

	return NormalizeEvents(all, namespace, pod, nil)
}

func unwrapRuntimeEvent(raw json.RawMessage) (RawRREvent, error) {
	// RR wraps Tetragon GetEventsResponse; tolerate flat or nested shapes.
	var flat RawRREvent
	if err := json.Unmarshal(raw, &flat); err == nil && flat.Process != nil {
		return flat, nil
	}
	var wrap struct {
		Event json.RawMessage `json:"event"`
		Time  string          `json:"registered_at"`
		Threats []RawThreat   `json:"threats"`
	}
	if err := json.Unmarshal(raw, &wrap); err != nil {
		return RawRREvent{}, err
	}
	var inner struct {
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
	_ = json.Unmarshal(wrap.Event, &inner)
	ev := RawRREvent{Time: wrap.Time, Threats: wrap.Threats}
	switch {
	case inner.ProcessExec != nil:
		ev.Type = "PROCESS_EXEC"
		ev.Process = inner.ProcessExec.Process
		ev.Parent = inner.ProcessExec.Parent
	case inner.ProcessExit != nil:
		ev.Type = "PROCESS_EXIT"
		ev.Process = inner.ProcessExit.Process
		ev.Parent = inner.ProcessExit.Parent
	case inner.ProcessKprobe != nil:
		ev.Type = "PROCESS_KPROBE"
		ev.Process = inner.ProcessKprobe.Process
		ev.Parent = inner.ProcessKprobe.Parent
		ev.FunctionName = inner.ProcessKprobe.FunctionName
		ev.Args = inner.ProcessKprobe.Args
	default:
		return RawRREvent{}, fmt.Errorf("unknown event shape")
	}
	return ev, nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
