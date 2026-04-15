---
acceptance: Calling EnsureRunning from a test starts syded when it's not running and is a no-op when it is
completed_at: "2026-04-14T06:45:25Z"
created_at: "2026-04-14T06:38:16Z"
details: 'NEW FILE internal/daemon/daemon.go: EnsureRunning tries GET /health with short timeout; on failure, forks syded with os.StartProcess; polls /health every 50ms up to 3s; returns error if still down. Uses pidfile at ~/.syde/syded.pid.'
id: TSK-0017
kind: task
name: Create internal/daemon package with EnsureRunning
objective: CLI can call EnsureRunning(port) and assume syded is up afterwards
plan_phase: phase_2
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: create-internaldaemon-package-with-ensurerunning-6fy5
task_status: completed
updated_at: "2026-04-14T06:45:25Z"
---
