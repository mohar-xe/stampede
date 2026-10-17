// Phase 0 — PROOF OF LIFE. This should run before anything else exists.
//
//   cd agent && bun install
//   export E2B_API_KEY=...   # e2b.dev dashboard, free, no card
//   bun run src/sandbox.ts
//
// Expected output: a fresh Linux VM boots, runs commands, clones a repo, dies.
// That's the entire sandbox primitive every YC agent company builds on.
//
// Docs: https://e2b.dev/docs · full index: https://docs.e2b.dev/llms.txt
// There's a public MCP docs server at https://docs.e2b.dev/mcp — ask it anything.

import { Sandbox } from "e2b";

const sb = await Sandbox.create(); // boots an isolated VM

try {
  // 1. Plain commands
  const uname = await sb.commands.run("uname -a");
  console.log("[kernel]", uname.stdout.trim());

  // 2. Real repo in a real filesystem
  await sb.commands.run(
    "git clone --depth 1 https://github.com/sindresorhus/is-odd /home/user/repo",
    { timeoutMs: 60_000 },
  );
  const files = await sb.commands.run("ls /home/user/repo | head");
  console.log("[repo files]\n" + files.stdout);

  // TODO(you): write a file, run its tests (`npm t`), read the output.
  // That loop — edit → test → observe — is 80% of what a coding agent is.
} finally {
  await sb.kill(); // sandboxes cost free-tier credits while alive — always kill
}
