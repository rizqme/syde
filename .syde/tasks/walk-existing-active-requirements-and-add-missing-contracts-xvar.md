---
id: TSK-0120
kind: task
name: Walk existing active requirements and add missing contracts
slug: walk-existing-active-requirements-and-add-missing-contracts-xvar
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-detector-coverage-symmetry-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: Every active requirement's surface has a corresponding active contract; sync check reports zero surface-coverage errors
details: Run syde sync check, collect surface-coverage errors; for each, read the requirement, decide whether to add a new contract (most common) or reword the requirement (when the surface was accidental phrasing). Use the existing contract-authoring flags including --contract-kind, --interaction-pattern, --input, --input-parameter, --output, --output-parameter. Where multiple requirements share a surface, one contract covers all of them.
acceptance: syde sync check reports zero surface-coverage errors; contract count grows to cover every existing requirement surface
affected_entities:
    - audit-engine-4ktg
    - cli-commands-hpjb
affected_files:
    - internal/audit/surfaces.go
    - internal/audit/surfaces_test.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_9
created_at: "2026-04-17T09:46:36Z"
completed_at: "2026-04-17T10:41:29Z"
---
