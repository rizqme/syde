---
acceptance: client.New works from a shell with no syded running — returns after syded is ready
completed_at: "2026-04-14T06:46:49Z"
created_at: "2026-04-14T06:38:16Z"
details: 'client.New calls daemon.EnsureRunning first. Error path includes actionable message (''syded failed to start: ...'').'
id: TSK-0021
kind: task
name: Wire daemon.EnsureRunning into client.New
objective: Client transparently auto-launches syded before the first request
plan_phase: phase_3
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: wire-daemonensurerunning-into-clientnew-g07d
task_status: completed
updated_at: "2026-04-14T06:46:49Z"
---
