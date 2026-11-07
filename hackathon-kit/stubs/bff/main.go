// STUB Go BFF — реализовать после старта в apps/bff.
// GET /api/activity → fixture или Runtime Radar public-api
// Секреты только из env (RR_BASE_URL, RR_TOKEN / RR_USER+RR_PASSWORD).
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("GET /api/activity", activityHandler)

	addr := envOr("ADDR", ":8000")
	log.Printf("activity-bff stub listening on %s (fixture-first)", addr)
	log.Fatal(http.ListenAndServe(addr, withCORS(mux)))
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("pod")
	_ = r.URL.Query().Get("namespace")

	data, src, err := loadActivity()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Data-Source", src)
	_ = json.NewEncoder(w).Encode(data)
}

func loadActivity() (any, string, error) {
	// TODO after start: try live RR, on failure fall back to fixture
	path := envOr("FIXTURE_PATH", filepath.Join("..", "..", "fixtures", "sample-activity.json"))
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	var data any
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, "", err
	}
	return data, "fixture", nil
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := envOr("CORS_ORIGIN", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
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
