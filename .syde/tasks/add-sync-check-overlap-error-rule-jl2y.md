---
id: TSK-0090
kind: task
name: Add sync check overlap error rule
slug: add-sync-check-overlap-error-rule-jl2y
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-audit-and-overlap-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: syde sync check ERRORs on unaudited requirement overlaps
details: New audit rule in internal/audit. For each pair of active requirements with >50% overlap, check both sides have each other in audited_overlaps. ERROR if not.
acceptance: Two similar requirements without audited_overlaps trigger ERROR; adding --audited clears it
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/plan_authoring.go
    - internal/audit/requirements.go
    - internal/audit/audit.go
plan_ref: requirement-overlap-audit-with-mandatory-acknowledgement
plan_phase: phase_2
created_at: "2026-04-16T11:29:52Z"
completed_at: "2026-04-16T11:43:58Z"
---
