---
acceptance: syde sync check --format json returns valid JSON with grouped findings
affected_files:
    - internal/cli/validate.go
completed_at: "2026-04-14T06:21:39Z"
created_at: "2026-04-14T06:04:02Z"
details: Payload {errors:[], warnings:[], hints:[]} using Finding from internal/audit. Enable --format json on validate and sync check.
id: TSK-0012
kind: task
name: JSON output for validate + sync check
objective: Health commands emit machine-readable findings
plan_phase: phase_3
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: json-output-for-validate-sync-check-0xd3
task_status: completed
updated_at: "2026-04-14T06:21:39Z"
---
