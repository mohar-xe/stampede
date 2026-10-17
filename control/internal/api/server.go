// Package api holds the HTTP surface of the control plane.
//
// The API mirrors Anthropic Managed Agents' surface (platform.claude.com/docs),
// deliberately simplified. Copy their concepts, not their bulk:
//
//	Agents      — reusable config: model, system prompt, tools      (POST /v1/agents)
//	Environments— sandbox template: image, setup script, env        (POST /v1/environments)
//	Sessions    — a running instance doing one task                 (POST /v1/sessions)
//	Events      — the session's append-only log                     (GET /v1/sessions/{id}/events)
//
// Phase 2 adds SSE streaming: GET /v1/sessions/{id}/events/stream
package api

import (
	"encoding/json"
	"net/http"
)

// Server owns the HTTP handlers. It depends on a Store interface, never on
// Postgres directly — that's what makes the session/harness/sandbox swap testable.
type Server struct {
	// TODO(you): store session.Store
	// TODO(you): queue queue.Queue
	mux *http.ServeMux
}

func New(/* store session.Store */) *Server {
	s := &Server{mux: http.NewServeMux()}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /healthz", s.handleHealth)

	// TODO(you) Phase 2:
	// s.mux.HandleFunc("POST /v1/agents", s.handleCreateAgent)
	// s.mux.HandleFunc("POST /v1/sessions", s.handleCreateSession)
	// s.mux.HandleFunc("GET /v1/sessions", s.handleListSessions)          // fleet view
	// s.mux.HandleFunc("GET /v1/sessions/{id}", s.handleGetSession)
	// s.mux.HandleFunc("GET /v1/sessions/{id}/events", s.handleListEvents)
	// s.mux.HandleFunc("GET /v1/sessions/{id}/events/stream", s.handleStream) // SSE
	// s.mux.HandleFunc("POST /v1/sessions/{id}/approve", s.handleApprove)     // ship gate
}

func (s *Server) ListenAndServe() error { return http.ListenAndServe(":8080", s.mux) }

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// TODO(you): implement handlers above. Event taxonomy suggestion (copied from
// Anthropic's webhooks): session.status_running / status_blocked / status_idled /
// status_terminated, plus harness.* for tool calls and usage.* for tokens/cost.

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
