---
id: TSK-0092
kind: task
name: Add BaseSlug fallback to handlePlanDetail task resolution
slug: add-baseslug-fallback-to-handleplandetail-task-resolution-b8mc
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-audit-and-overlap-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: Tasks resolve by bare slug when exact match fails
acceptance: Plan detail API returns task_index with all tasks including 'build-and-verify'
affected_entities:
    - http-api-afos
affected_files:
    - internal/dashboard/api.go
plan_ref: fix-task-resolution-in-plan-detail-api
plan_phase: phase_1
created_at: "2026-04-16T11:49:56Z"
completed_at: "2026-04-16T11:52:04Z"
---
