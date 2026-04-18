---
id: TSK-0247
kind: task
name: Add --refined-by flag to syde query
slug: add-refined-by-flag-to-syde-query-gsg3
updated_at: '2026-04-18T08:17:51Z'
task_status: completed
priority: medium
objective: syde query --refined-by <component-slug> lists active reqs refining the named component, using the existing rich/json/compact/refs formatters
acceptance: syde query --refined-by audit-engine-4ktg returns refining reqs; --format refs returns plain slug list
affected_files:
- internal/cli/query.go
plan_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
plan_phase: phase_4
created_at: '2026-04-18T08:00:05Z'
completed_at: '2026-04-18T08:17:51Z'
relationships:
- type: belongs_to
  target: syde-5tdt
- type: implements
  target: syde-query-shall-support-refined-by-component-slug-o23d
---
