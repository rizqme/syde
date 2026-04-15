---
acceptance: Start syded with --idle-timeout 2s; wait 5s; process has exited and pidfile removed
affected_entities:
    - syded-dashboard
affected_files:
    - internal/dashboard/run.go
completed_at: "2026-04-14T06:44:16Z"
created_at: "2026-04-14T06:38:16Z"
details: 'Middleware records last-request timestamp. Ticker every 60s: if elapsed > --idle-timeout, call srv.Shutdown() and remove pidfile. --idle-timeout flag, default 30m, 0 disables.'
id: TSK-0019
kind: task
name: Add idle-timeout auto-shutdown to syded
objective: syded dies cleanly after 30m of no /api traffic so it doesn't linger forever
plan_phase: phase_2
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-idle-timeout-auto-shutdown-to-syded-vlnr
task_status: completed
updated_at: "2026-04-14T06:44:16Z"
---
