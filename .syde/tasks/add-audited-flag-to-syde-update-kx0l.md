---
id: TSK-0089
kind: task
name: Add --audited flag to syde update
slug: add-audited-flag-to-syde-update-kx0l
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-audit-and-overlap-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: syde update <req> --audited <slug> appends to audited_overlaps
details: Add flag to update.go requirement case; append to existing AuditedOverlaps
acceptance: syde update <req> --audited <slug> adds to the list
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/update.go
plan_ref: requirement-overlap-audit-with-mandatory-acknowledgement
plan_phase: phase_1
created_at: "2026-04-16T11:29:52Z"
completed_at: "2026-04-16T11:42:19Z"
---
