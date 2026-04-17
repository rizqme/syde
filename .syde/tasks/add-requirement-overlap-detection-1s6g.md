---
id: TSK-0075
kind: task
name: Add requirement overlap detection
slug: add-requirement-overlap-detection-1s6g
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-plan-requirement-coverage-and-overlap-audit
      type: references
updated_at: "2026-04-16T09:48:49Z"
task_status: completed
objective: WARN when a new requirement overlaps an existing active one by term similarity
details: For each new requirement draft statement, tokenize into significant words (>3 chars, lowered, stop words removed). Compare against all existing active requirement statements. If >50% shared terms, emit WARN naming both.
acceptance: syde plan check on a plan with a near-duplicate requirement shows the overlap warning
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/plan_authoring.go
plan_ref: plan-requirement-coverage-and-overlap-audit
plan_phase: phase_1
created_at: "2026-04-16T09:44:03Z"
completed_at: "2026-04-16T09:48:22Z"
---
