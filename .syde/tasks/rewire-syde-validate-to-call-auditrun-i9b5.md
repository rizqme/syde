---
acceptance: syde validate errors on a planted orphan; still errors on cycle; passes on clean repo
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/validate.go
completed_at: "2026-04-14T06:14:07Z"
created_at: "2026-04-14T06:03:22Z"
details: Replace validate.go body with audit.Run + grouped print + exit 1 on any Error. Keep existing cycle + missing-field checks (route them through audit categories).
id: TSK-0001
kind: task
name: Rewire syde validate to call audit.Run
objective: syde validate now fails on orphan files
plan_phase: phase_1
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: rewire-syde-validate-to-call-auditrun-i9b5
task_status: completed
updated_at: "2026-04-14T06:14:07Z"
---
