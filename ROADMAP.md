# Stampede Roadmap

Each phase ships something demoable. Every phase has a **learning goal** — the infra
concept it teaches. Implementation is left to you; the stubs + TODOs are the trailheads.

## Phase 0 — Proof of life (sandbox)          ← start here
- [ ] `cd agent && bun install`
- [ ] Set `E2B_API_KEY`
- [ ] `bun run src/sandbox.ts` → creates an E2B VM, runs commands, clones a repo
- **Learning goal**: what a code sandbox actually is (VM vs container vs isolate),
  why every YC agent company uses one.

## Phase 1 — The harness
- [ ] Implement the loop in `agent/src/harness.ts`:
      plan → emit file edits → run tests in sandbox → observe → repeat (max N turns)
- [ ] Streaming events printed to stdout (SSE-shaped JSON lines)
- **Learning goal**: Anthropic's "Building Effective Agents" — simple loop + tools,
  ground-truth from the environment, error recovery. No framework needed.

## Phase 2 — Session log + control plane
- [ ] Postgres schema (`control/internal/store` TODO) via Neon
- [ ] `POST /v1/sessions`, `GET /v1/sessions/{id}/events`, `GET .../events/stream` (SSE)
- [ ] Harness writes events to the session log; dashboard reads it
- **Learning goal**: session-as-append-only-log (Anthropic's core insight),
  REST + SSE streaming, idempotency.

## Phase 3 — Fleet + queue
- [ ] Job queue so K harnesses run in parallel on K sessions
- [ ] Fleet view: states (queued/running/blocked/done/failed), usage & cost per session
- **Learning goal**: worker queues, concurrency limits, observability.
  "When an agent fails quietly, the trace is the product."

## Phase 4 — Ship gate
- [ ] Sandbox pushes branch → GitHub PR via API
- [ ] Human approval step in the dashboard before PR creation
- **Learning goal**: guardrails/permissions (Runtime's whole pitch), git plumbing.

## Phase 5 — Deploy + polish
- [ ] Go control plane on Render free; Postgres → Neon; dashboard → Vercel
- [ ] README demo GIF; write the HN-style launch post
- **Learning goal**: free-tier deployment end-to-end, spinning down, env management.

## Stretch (pick what's fun)
- Slack/webhook triggers (`session.status_*` events — copy Anthropic's taxonomy)
- "Group chat" UX like Agent FM (agents push blockers to a chat)
- Shared memory across sessions (Mosaic-style)
- Swap E2B for Cloudflare Containers later if you ever want to own that layer
