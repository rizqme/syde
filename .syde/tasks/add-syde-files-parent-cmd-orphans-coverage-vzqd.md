---
acceptance: syde files orphans prints nothing on clean repo; syde files coverage internal/cli/add.go prints owning component
affected_entities:
    - cli-commands
completed_at: "2026-04-14T06:15:08Z"
created_at: "2026-04-14T06:04:02Z"
details: 'NEW FILE internal/cli/files.go: cobra parent ''files'' + ''files orphans'' (one path per line, exit 1 if any) + ''files coverage [path]''. Register in init(). Reuses internal/audit.'
id: TSK-0009
kind: task
name: Add syde files parent cmd + orphans + coverage
objective: Targeted orphan + coverage listings without running full validate
plan_phase: phase_1
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-syde-files-parent-cmd-orphans-coverage-vzqd
task_status: completed
updated_at: "2026-04-14T06:15:08Z"
---
