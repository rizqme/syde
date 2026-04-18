---
id: TSK-0267
kind: task
name: Implement 'syde plan review <slug>' CLI command
slug: implement-syde-plan-review-slug-cli-command-ax3u
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: syde-cli-shall-provide-plan-review-command-that-dispatches-a-plan-reviewer-subagent-6wli
      type: implements
updated_at: "2026-04-18T09:44:54Z"
task_status: completed
priority: medium
objective: syde plan review <slug> loads the plan markdown, loads the reviewer prompt from the embedded skill bundle, prints a ready-to-paste reviewer prompt (or invokes it if a subagent backend is configured) and returns the reviewer's Approved or Issues Found verdict.
details: Add 'review <slug>' under planCmd in internal/cli/plan.go. Resolve plan via store.Get; read the plan markdown bytes. Load the reviewer prompt from skill/references/plan-review-prompt.md (new file, bundled via go:embed in internal/skill/embed.go). In the v1 implementation, print the reviewer prompt with the plan content interpolated, and let the user paste it into a subagent themselves. A later version can call a local LLM backend or the Claude API directly.
acceptance: syde plan review <existing-plan-slug> prints a reviewer prompt with the plan content embedded; exits 0 regardless of reviewer verdict (verdict is in the output, not the exit code).
affected_entities:
    - cli-commands-hpjb
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_4
created_at: "2026-04-18T09:24:42Z"
completed_at: "2026-04-18T09:36:10Z"
---
