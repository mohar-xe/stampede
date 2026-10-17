// Phase 1 — the harness: the loop that calls the LLM and routes tool calls.
//
// Design rules (from Anthropic — read both before implementing):
//   1. https://www.anthropic.com/engineering/building-effective-agents
//      → keep it a SIMPLE loop + tools; complexity is where agents break.
//   2. https://www.anthropic.com/engineering/managed-agents
//      → the harness is swappable; it talks only to the session log + sandbox.
//
// No agent framework. Bun + an LLM SDK + E2B is the whole dependency list.

// ---------------------------------------------------------------------------
// Events — the harness's only output. Everything else is an implementation.
// These shape the session log the control plane stores and the dashboard plays back.
// ---------------------------------------------------------------------------
export type HarnessEvent =
  | { type: "harness.plan"; steps: string[] }
  | { type: "harness.thought"; text: string }
  | { type: "harness.tool_call"; tool: string; args: unknown }
  | { type: "harness.tool_result"; tool: string; ok: boolean; output: string }
  | { type: "harness.done"; summary: string }
  | { type: "harness.error"; error: string }
  | { type: "usage.tokens"; input: number; output: number; usd: number };

type Emit = (e: HarnessEvent) => void;

// For Phase 1 just pretty-print JSON lines (NDJSON ≈ the session log format).
export const stdoutEmit: Emit = (e) =>
  console.log(JSON.stringify({ ts: Date.now(), ...e }));

// ---------------------------------------------------------------------------
// Tools — the verbs the model can use. Each one touches the sandbox.
// ---------------------------------------------------------------------------
export type Tool = {
  name: string;
  description: string;
  run: (sb: unknown, args: unknown) => Promise<string>; // TODO(you): real Sandbox type
};

// TODO(you) implement:
//  - run_command { cmd }          → sb.commands.run(cmd)
//  - read_file  { path }          → sb.files.read(path)
//  - write_file { path, content } → sb.files.write(path, content)
// These three are enough. Devin/Claude Code started with roughly this shape.

// ---------------------------------------------------------------------------
// The loop — plan → act → observe → repeat, with a turn budget.
// ---------------------------------------------------------------------------
export const MAX_TURNS = 15;

export async function runSession(
  task: string,
  emit: Emit = stdoutEmit,
): Promise<void> {
  emit({ type: "harness.thought", text: `task: ${task}` });

  // TODO(you):
  //  1. Create the sandbox.                          (agent/src/sandbox.ts shows how)
  //  2. System prompt: role, tools, "always run tests before claiming done".
  //  3. Loop MAX_TURNS times:
  //       - ask the model for ONE tool call (JSON mode / tool-use)
  //       - emit harness.tool_call → execute → emit harness.tool_result
  //       - feed the result back in
  //       - stop when the model calls a `done` tool or the budget is spent
  //  4. Kill the sandbox in a finally block.
  //
  // Error recovery is the learning goal: when a command fails, the model sees
  // stderr and retries — that's the whole magic. Watch it happen. Trace it.
  throw new Error("TODO(you): implement the loop — see ROADMAP.md Phase 1");
}
