---
id: TSK-0122
kind: task
name: Plan authoring warns when contract changes lack flow coverage
slug: plan-authoring-warns-when-contract-changes-lack-flow-coverage-0x2m
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-detector-coverage-symmetry-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: syde plan check emits WARN when a plan's contract lane has new or extended entries not covered by any flow new or extended entry whose steps reference the contract
details: 'Extend internal/audit/plan_authoring.go: for each new/extended contract change, compute its slug (for new, the declared name slugified; for extended, the target slug); scan the plan''s flow lane — new flows'' --draft steps field and extended flows'' field_changes on steps — for a step whose contract value matches. Emit WARN if gap with finding key ''contract_flow_coverage'' and message naming the uncovered contract.'
acceptance: syde plan check on a crafted plan with a new contract and no flow entry warns; plan with matching flow step does not
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/plan_authoring.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_11
created_at: "2026-04-17T09:50:47Z"
completed_at: "2026-04-17T10:23:23Z"
---
