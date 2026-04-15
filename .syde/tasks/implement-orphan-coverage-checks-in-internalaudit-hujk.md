---
acceptance: Clean repo returns 0 orphans; planting a stray .go file returns 1 orphan finding
completed_at: "2026-04-14T06:12:13Z"
created_at: "2026-04-14T06:04:02Z"
details: 'NEW FILE internal/audit/orphans.go: load .syde/tree.yaml via internal/tree.Load, collect non-ignored file nodes, load all entities via store.List, union entity.Files, diff. FileCoverage returns map[path][]ownerSlug. Handle duplicate owners.'
id: TSK-0008
kind: task
name: Implement orphan + coverage checks in internal/audit
objective: Orphan file detection works end-to-end against real repo data
plan_phase: phase_1
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: implement-orphan-coverage-checks-in-internalaudit-hujk
task_status: completed
updated_at: "2026-04-14T06:12:13Z"
---
