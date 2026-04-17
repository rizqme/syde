---
id: TSK-0088
kind: task
name: Add overlap detection to syde add requirement
slug: add-overlap-detection-to-syde-add-requirement-za03
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-audit-and-overlap-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: syde add requirement prints similar requirements and accepts --audited
details: In add.go requirement case, after store.Create, run significantTerms/termOverlap against all active requirements. Print matches. Wire --audited flag to populate AuditedOverlaps.
acceptance: Creating a near-duplicate requirement prints the overlap; --audited stores the slug
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/add.go
plan_ref: requirement-overlap-audit-with-mandatory-acknowledgement
plan_phase: phase_1
created_at: "2026-04-16T11:29:52Z"
completed_at: "2026-04-16T11:40:37Z"
---
