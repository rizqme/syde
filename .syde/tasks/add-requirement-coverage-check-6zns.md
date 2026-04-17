---
id: TSK-0074
kind: task
name: Add requirement coverage check
slug: add-requirement-coverage-check-6zns
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-plan-requirement-coverage-and-overlap-audit
      type: references
updated_at: "2026-04-16T09:48:49Z"
task_status: completed
objective: WARN when requirements < non-requirement changes / 3
details: Count non-req changes across all lanes. Count req changes. If ratio too low, emit WARN.
acceptance: syde plan check on a plan with 9 changes and 1 requirement shows the warning
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/plan_authoring.go
plan_ref: plan-requirement-coverage-and-overlap-audit
plan_phase: phase_1
created_at: "2026-04-16T09:44:03Z"
completed_at: "2026-04-16T09:48:22Z"
---
