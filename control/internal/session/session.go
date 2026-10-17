// Package session defines the session primitive: an append-only log.
//
// This is the single most important design decision in the project (Anthropic's
// engineering post "Scaling Managed Agents" is built on it):
//
//   - The session log is the source of truth. Nothing else is.
//   - Harnesses and sandboxes are swappable *because* they only touch the log.
//   - The dashboard is just a reader of the log. SSE = tail -f.
//
// Learn from: anthropic.com/engineering/managed-agents
// and OpenHands' event-stream architecture (github.com/all-hands-ai/openhands).
package session

import "time"

// Session is one agent run doing one task.
type Session struct {
	ID        string    `json:"id"`
	AgentID   string    `json:"agent_id"` // which reusable config runs this
	Task      string    `json:"task"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Status string

// Suggested lifecycle — mirror Anthropic's webhook taxonomy so the dashboard
// reads exactly like their docs.
const (
	StatusQueued     Status = "queued"
	StatusRunning    Status = "status_running"
	StatusBlocked    Status = "status_blocked" // waiting on human approval
	StatusIdled      Status = "status_idled"
	StatusTerminated Status = "status_terminated"
	StatusFailed     Status = "status_failed"
)

// Event is one immutable entry in the session log.
type Event struct {
	ID        int64     `json:"id"` // position in the log — never reorder, never rewrite
	SessionID string    `json:"session_id"`
	Type      string    `json:"type"`   // e.g. "harness.plan", "harness.tool_call", "usage.tokens"
	Payload   string    `json:"payload"` // raw JSON blob; don't over-model early
	CreatedAt time.Time `json:"created_at"`
}

// Store is the persistence boundary. Implement in-memory first, Postgres second.
// The interface IS the lesson: the log outlives every implementation.
type Store interface {
	CreateSession(s Session) error
	GetSession(id string) (Session, error)
	ListSessions() ([]Session, error)

	// Append must be crash-safe and order-preserving.
	AppendEvent(e Event) error
	Events(sessionID string, afterID int64) ([]Event, error)
}
