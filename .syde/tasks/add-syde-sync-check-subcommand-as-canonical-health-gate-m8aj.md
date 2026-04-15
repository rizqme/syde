---
acceptance: syde sync check --strict exits 0 clean; orphan makes it exit 1
affected_entities:
    - cli-commands
completed_at: "2026-04-14T06:16:50Z"
created_at: "2026-04-14T06:04:02Z"
details: NEW FILE internal/cli/sync.go (or promote existing sync --check flag). Subcommand 'sync check' calls audit.Run then tree.Status then drift. Prints Errors/Warnings/Hints with counts. Exit 0 clean / 1 error / 2 strict+warning.
id: TSK-0010
kind: task
name: Add syde sync check subcommand as canonical health gate
objective: One command runs audit + tree-status + drift and exits non-zero on any gap
plan_phase: phase_2
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-syde-sync-check-subcommand-as-canonical-health-gate-m8aj
task_status: completed
updated_at: "2026-04-14T06:16:50Z"
---
