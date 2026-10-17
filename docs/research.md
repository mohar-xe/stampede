# Research notes (Oct 2026)

Everything that shaped this project's design, with links.

## The competitive landscape (who we're cloning, at the intersection)

- **Runtime (YC P26)** — control plane for coding agents: sandboxes, guardrails,
  observability, sessions from Slack/Linear/CLI/browser. Open source.
  - https://github.com/runtm-ai/runtm
  - https://news.ycombinator.com/item?id=48225040 (launch thread, 81 comments)
  - https://www.runtm.com/pricing — their wedge: parallel sessions, approvals,
    usage tracking, env templates
- **herdr (YC F26)** — open-source CLI agent runtime, owns terminal panes,
  agents survive disconnects. 25k★ / 340k downloads as a *solo* project before YC.
  - https://herdr.dev
  - https://www.ycombinator.com/companies/herdr
  - https://herdr.dev/blog/herdr-is-joining-y-combinator (open-core Apache-2.0 story)
- **Agent FM (YC S26)** — "group chat with your Claude Code and Codex agents":
  agents push progress/blockers/decisions; live text+audio updates. Newest UX idea.
- **Mosaic (YC 2026)** — shared memory/context across coding-agent sessions;
  "defined the multiplayer AI category".
- **Claude Managed Agents (Anthropic, Apr 8 2026)** — Anthropic shipping the same
  control-plane layer as a product.

## The canonical architecture — read Anthropic's post twice

https://www.anthropic.com/engineering/managed-agents — "Decoupling the brain from the hands"

They virtualize an agent into three swappable components (this repo's primitives):
- **session** — the append-only log of everything that happened
- **harness** — the loop that calls Claude and routes tool calls to infra
- **sandbox** — the execution environment

Key quote: harnesses encode assumptions about what the model can't do on its own —
and those assumptions go stale. So: opinionated interfaces, not implementations.

Also by Anthropic:
- https://www.anthropic.com/engineering/building-effective-agents — the simple
  agentic loop pattern; ground-truth from the environment; progress visibility.
- 2026 Agentic Coding Trends Report (resources.anthropic.com) — long-running
  multi-hour agents, humans as orchestrators, fleet dashboards.
- Hookdeck on Managed Agents webhooks: `session.status_run_started`,
  `session.status_idled`, `session.status_terminated`, vault events — good event
  taxonomy to copy for our session log.
  https://hookdeck.com/blog/anthropic-managed-agent-webhooks

## YC batch data

- https://www.tldl.io/blog/yc-ai-startups-2026 — 2026 batches: agent reliability
  is becoming its own stack (Arga Labs: sandbox testing; Syntropy W26: agents
  running generated tests; Salus: validate-before-execute).
- https://e2b.dev/resources/yc-companies-ai-agents — ~50% of recent YC batches are
  agent companies; most outsource sandbox infra to E2B.
- Runtime's own words: "letting more people across your org use coding agents
  without creating chaos" = the guardrails/observability gap = our product.

## Sandbox infra (our choice: E2B)

- **E2B** — VM-per-sandbox, JS/Python SDK, free tier without card.
  - https://e2b.dev/docs · full index: https://docs.e2b.dev/llms.txt
  - MCP docs server: https://docs.e2b.dev/mcp (no auth) — ask it questions
- **Cloudflare Sandbox SDK** — great, BUT Containers require Workers Paid ($5/mo
  + card) → ruled out for us. Bookmark for later.
  - https://developers.cloudflare.com/sandbox
  - https://github.com/cloudflare/sandbox-sdk
- Sanity check on CF sandbox pricing: containers bill per vCPU-second on the
  paid plan — fine for a company, not for $0.

## Reference open source to study (read, don't copy)

- runtm/runtm — FastAPI control plane + Go CLI + sandbox packages
- herdr — terminal-pane runtime, TUI, machine management
- OpenHands (all-hands-ai) — event-stream architecture, very close to
  session-as-log
- cloudflare/sandbox-sdk — Durable Object per sandbox pattern
