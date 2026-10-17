// Package queue is the job queue that turns one session into one harness run.
//
// Phase 3 learning goal: why a queue at all?
//   - Decouple API latency from agent runtime (a session runs for *hours*).
//   - Bounded parallelism (K workers = K sandboxes = $0 stays $0).
//   - Retries + visibility when a harness dies mid-run.
//
// Implementation hints, in rising order of learning value:
//   1. Postgres-backed queue: SELECT ... FOR UPDATE SKIP LOCKED — one query,
//      no extra infra, works on Neon free. (Do this.)
//   2. In-process channel queue — fine for local dev only.
//   3. Read how Render/Fly/Cloudflare Queues frame at-least-once delivery and
//      why Anthropic tells webhook consumers "dedupe upstream or pay twice"
//      (hookdeck.com/blog/anthropic-managed-agent-webhooks).
package queue

import "context"

// Job = "run session S in sandbox, with harness H, write events back".
// Keep it a dumb struct: the interesting logic lives in the worker, not the queue.
type Job struct {
	SessionID string
	Attempt   int // increment on retry; after N attempts → session failed
}

type Queue interface {
	Enqueue(ctx context.Context, j Job) error
	// Dequeue blocks until a job is available or ctx is done.
	Dequeue(ctx context.Context) (Job, error)
	// Ack/Nack: at-least-once delivery means the worker must confirm.
	Ack(ctx context.Context, j Job) error
	Nack(ctx context.Context, j Job, reason string) error
}
