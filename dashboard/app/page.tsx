// Fleet view — Phase 3 target: list all sessions + live status.
// Phase 2 target: list events of ONE session (a replayable trace).
//
// Data flow: this page hits the Go control plane (SSE for live updates).
// TODO(you): keep the control-plane URL in an env var —
//   NEXT_PUBLIC_CONTROL_PLANE_URL=http://localhost:8080 (dev) → Render URL (prod)
export default function FleetPage() {
  // TODO(you) Phase 2: fetch("/v1/sessions") and render a table:
  //   id · task · status · tokens · usd · elapsed
  // TODO(you) Phase 2: clicking a row opens /sessions/[id] — the event log,
  //   rendered like a terminal. "When an agent fails quietly, the trace is
  //   the product" — make the trace the hero of this dashboard.
  // TODO(you) Phase 3: EventSource("/v1/sessions/stream") for live status.
  // TODO(you) Phase 4: "Approve → ship PR" button (POST /v1/sessions/{id}/approve).
  return (
    <main style={{ fontFamily: "monospace", padding: 32 }}>
      <h1>stampede · fleet</h1>
      <p style={{ color: "#888" }}>
        No sessions yet. Implement the control plane, then come back here.
      </p>
    </main>
  );
}
