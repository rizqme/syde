---
acceptance: Building the old corrupted plan's data shape (5 phases, all IDs='') through CollectTasks no longer infinite-loops. syde plan show on a plan with a deliberate self-parent does not crash.
affected_entities:
    - entity-model
affected_files:
    - internal/model/plan.go
completed_at: "2026-04-14T08:57:09Z"
created_at: "2026-04-14T08:55:51Z"
details: 'In internal/model/plan.go: (1) ChildPhases — skip ph.ID == '''' and ph.ID == parentID. (2) CollectTasks — split into public wrapper that allocates a visited map[string]bool, and an internal helper (phaseID, visited) that returns early if visited[phaseID]. Mark visited on entry, recurse on each child. Preserves existing callers.'
id: TSK-0048
kind: task
name: Guard ChildPhases and CollectTasks against cycles and empty IDs
objective: PlanEntity.CollectTasks returns cleanly on any data shape (self-loops, empty IDs, ParentPhase cycles)
plan_phase: phase_1
plan_ref: guard-plan-phase-integrity
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: guard-childphases-and-collecttasks-against-cycles-and-empty-ids-ohp6
task_status: completed
updated_at: "2026-04-14T08:57:09Z"
---
