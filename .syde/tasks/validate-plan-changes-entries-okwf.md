---
id: TSK-0002
kind: task
name: Validate plan changes entries
slug: validate-plan-changes-entries-okwf
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: ValidateEntity rejects a plan whose Changes entries are missing what/why, or whose NewChange Draft is empty.
details: 'Extend validation.go PlanEntity case: iterate all six ChangeLane fields; for Extended and New require non-empty What+Why; for Deleted require Why; for NewChange require non-empty Name and Draft with at least the kind-required fields (component must have responsibility+≥1 capability, contract must have input+output+contract_kind, requirement must match EARS via model.MatchEARS on Draft[''statement''], etc.).'
acceptance: 'Unit-equivalent: syde add/update plan with an empty-Why change fails validation.'
affected_entities:
    - entity-model-f28o
affected_files:
    - internal/model/validation.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_1
created_at: "2026-04-15T11:40:36Z"
completed_at: "2026-04-15T11:47:47Z"
---
