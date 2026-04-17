---
id: TSK-0046
kind: task
name: Implement syde plan open CLI command
slug: implement-syde-plan-open-cli-command-t7bx
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-approvals-shall-be-preceded-by-a-dashboard-open-cmz5
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: syde plan open <plan> resolves the plan slug, sends POST /api/<proj>/navigate, and if no clients are listening falls back to spawning a new browser tab via OS open (macOS) / xdg-open (linux) / start (windows).
details: 'internal/cli/plan.go: new planOpenCmd. Resolve project slug from the daemon registry. Send POST navigate; check response client count. If clients > 0, print ''opened in existing dashboard tab''. If clients == 0, spawn os.exec(open <url>) (cross-platform helper). Print the URL either way.'
acceptance: syde plan open plans-inbox-2-column-layout switches an open dashboard tab to the plan detail page; if no tab is open, it spawns one.
affected_entities:
    - cli-commands-hpjb
    - cli-http-client-otp2
affected_files:
    - internal/cli/plan.go
    - internal/client/client.go
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:26:48Z"
completed_at: "2026-04-15T15:12:36Z"
---
