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

// fetchLiveActivity — Damir live wiring.
// History from "now" goes BACKWARD via /slice/right (registered_at < cursor).
// /slice/left is futureward and wrong for completed gen-activity dumps.
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
	direction := strings.ToLower(envOr("RR_SLICE_DIR", "right")) // right = history

	client := &http.Client{Timeout: 30 * time.Second}
	if os.Getenv("RR_INSECURE_TLS") == "true" || os.Getenv("RR_INSECURE_TLS") == "1" {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
		}
	}

	var all []RawRREvent
	var skipped int
	cursor := time.Now().UTC().Format(time.RFC3339Nano)

	for i := 0; i < pages; i++ {
		u, err := url.Parse(fmt.Sprintf("%s/api/v1/public-api/runtime-event/slice/%s", base, direction))
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
			RightCursor   string            `json:"right_cursor"`
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
				skipped++
				continue
			}
			all = append(all, ev)
		}
		next := page.RightCursor
		if direction == "left" {
			next = page.LeftCursor
		}
		if next == "" || next == cursor {
			break
		}
		cursor = next
	}

	if skipped > 0 {
		fmt.Printf("live: skipped %d undecodable events (kept %d)\n", skipped, len(all))
	}
	return NormalizeEvents(all, namespace, pod, nil)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
