---
id: TSK-0053
kind: task
name: Fix phase auto-complete task matching
slug: fix-phase-auto-complete-task-matching-o44f
relationships:
    - target: syde
      type: belongs_to
    - target: approved-plan-fix-phase-auto-completion-cross-plan-collision
      type: references
updated_at: "2026-04-16T08:15:52Z"
task_status: completed
objective: setTaskStatus matches tasks by BaseSlug and scopes to same plan
details: In task.go lines 574-582, replace utils.Slugify(te.Name) == tSlug with utils.BaseSlug(te.GetBase().Slug) == tSlug, and add te.PlanRef == t.PlanRef
acceptance: go build clean; phase auto-completes correctly when duplicate task names exist across plans
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/task.go
plan_ref: fix-phase-auto-completion-cross-plan-collision
plan_phase: phase_1
created_at: "2026-04-16T05:12:14Z"
completed_at: "2026-04-16T05:16:21Z"
---
