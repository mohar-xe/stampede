# stampede 🐎

An open-source **control plane for AI coding-agent fleets** — the infrastructure layer
that YC is currently funding (Runtime YC P26, herdr YC F26, Agent FM YC S26) and that
Anthropic itself shipped as "Claude Managed Agents" (Apr 2026).

You run N coding-agent sessions in parallel, each in its own cloud sandbox. The control
plane owns the session log, streams live progress, tracks usage/cost, and gates shipping
(approve → PR).

> Name is a placeholder — rename freely.

## Architecture (Anthropic's own decomposition, from their engineering blog)

```
┌────────────────────────────────────────────────────────────┐
│  dashboard (Next.js on Vercel free)                        │
│  fleet view · live session feed · diffs · approvals        │
└──────────────────────────┬─────────────────────────────────┘
                           │ REST + SSE
┌──────────────────────────▼─────────────────────────────────┐
│  control plane (Go on Render free)                         │
│  sessions · agents · templates · usage · job queue         │
│  "session = append-only log of everything that happened"   │
└──────────┬───────────────────────────────┬─────────────────┘
           │ dispatch jobs                 │ stream events
┌──────────▼───────────────┐   ┌───────────▼─────────────────┐
│  harness (TypeScript/bun)│   │  sandboxes (E2B free tier)  │
│  the agent loop: plan →  │──▶│  one isolated VM per task:  │
│  edit → test → iterate   │   │  clone repo, run commands   │
└──────────────────────────┘   └─────────────────────────────┘
```

Three swappable primitives (keep these interfaces clean — that's the whole design lesson):

| Primitive | Lives in | Meaning |
|---|---|---|
| **session** | control plane | Append-only event log of a run. Source of truth. |
| **harness** | `agent/` | The loop that calls the LLM and routes tool calls. |
| **sandbox** | E2B (swap later for Docker/local) | Where code actually executes. |

## 100% free deployment (no credit card anywhere)

| Piece | Service | Why |
|---|---|---|
| Dashboard | Vercel free | Next.js native |
| Control plane | Render free web service | Go works, spins down when idle |
| Postgres | Neon free | Serverless PG, no card |
| Sandboxes | E2B free tier | VM-per-task, no card, API key from dashboard |
| Repos/PRs | GitHub free | fine-grained PAT |
| LLM | your choice (Claude free credits, Gemini free tier, …) | keep provider-agnostic in the harness |

## Accounts checklist (5 min, no card)

- [ ] e2b.dev → dashboard → copy `E2B_API_KEY`
- [ ] neon.tech → new project → copy `DATABASE_URL`
- [ ] render.com → (later, at deploy phase)
- [ ] vercel.com → (later, at deploy phase)
- [ ] github.com → fine-grained token with repo scope

## Layout

```
control/    Go control plane — REST API, session store, job queue   (you implement)
agent/      TS harness — the agent loop + E2B sandbox scripts      (you implement)
dashboard/  Next.js fleet view                                     (you implement)
docs/       research notes + all reference links
```

Start with [`docs/research.md`](docs/research.md), then [`ROADMAP.md`](ROADMAP.md).
Every `// TODO` in the code is yours.
