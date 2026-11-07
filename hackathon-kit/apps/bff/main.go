// Activity Map BFF (Go) — OnlyGo apps/bff.
// Fixture-first; live via REST public-api slice (Damir wiring).
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("GET /api/activity", activityHandler)

	addr := envOr("ADDR", ":8000")
	log.Printf("activity-bff listening on %s (DATA_MODE=%s)", addr, envOr("DATA_MODE", "fixture"))
	log.Fatal(http.ListenAndServe(addr, withCORS(mux)))
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	pod := r.URL.Query().Get("pod")
	ns := r.URL.Query().Get("namespace")
	if pod == "" {
		pod = "attacker-pod"
	}
	if ns == "" {
		ns = "hackathon"
	}

	payload, src, err := loadActivity(ns, pod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if payload.Pod == "" {
		payload.Pod = pod
	}
	if payload.Namespace == "" {
		payload.Namespace = ns
	}
	ensureSlices(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Data-Source", src)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	_ = enc.Encode(payload)
}

func loadActivity(namespace, pod string) (*ActivityPayload, string, error) {
	mode := strings.ToLower(envOr("DATA_MODE", "fixture"))
	fixturePath := envOr("FIXTURE_PATH", "fixtures/sample-activity.json")
	rawPath := envOr("RAW_DUMP_PATH", "")

	switch mode {
	case "live":
		p, err := fetchLiveActivity(namespace, pod)
		if err != nil {
			log.Printf("live failed (%v) → fixture", err)
			p2, err2 := loadFixture(fixturePath)
			if err2 != nil {
				return nil, "", err2
			}
			return p2, "fixture", nil
		}
		return p, "live", nil
	case "raw":
		if rawPath == "" {
			rawPath = "../../fixtures/raw/events.json"
		}
		p, err := NormalizeRawFile(rawPath, namespace, pod, nil)
		if err != nil {
			log.Printf("raw failed (%v) → fixture", err)
			p2, err2 := loadFixture(fixturePath)
			if err2 != nil {
				return nil, "", err2
			}
			return p2, "fixture", nil
		}
		return p, "raw", nil
	default:
		p, err := loadFixture(fixturePath)
		if err != nil {
			return nil, "", err
		}
		return p, "fixture", nil
	}
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := envOr("CORS_ORIGIN", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Access-Token")
		w.Header().Set("Access-Control-Expose-Headers", "X-Data-Source")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
