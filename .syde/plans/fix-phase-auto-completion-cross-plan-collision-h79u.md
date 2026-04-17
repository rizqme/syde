---
id: PLN-0003
kind: plan
name: Fix phase auto-completion cross-plan collision
slug: fix-phase-auto-completion-cross-plan-collision-h79u
relationships:
    - target: approved-plan-fix-phase-auto-completion-cross-plan-collision-q9u9
      type: references
      label: requirement
    - target: syde
      type: belongs_to
updated_at: "2026-04-16T08:16:06Z"
plan_status: completed
background: Phase auto-completion in setTaskStatus matches tasks by utils.Slugify(te.Name) without scoping to the current plan. When two tasks from different plans share the same slugified name and one is pending, the check falsely concludes not all tasks are done, blocking auto-completion.
objective: Phase auto-complete correctly scopes task lookup to the same plan and matches by entity slug instead of re-slugifying the name.
scope: Fix the inner loop in setTaskStatus (internal/cli/task.go lines 574-582). No other changes needed.
design: Replace utils.Slugify(te.Name) == tSlug with utils.BaseSlug(te.GetBase().Slug) == tSlug, and add te.PlanRef == t.PlanRef to scope the match to the same plan.
source: manual
created_at: "2026-04-16T05:11:31Z"
approved_at: "2026-04-16T05:15:50Z"
completed_at: "2026-04-16T08:16:06Z"
phases:
    - id: phase_1
      name: Fix and verify
      status: completed
      description: Fix the slug matching and plan scoping in setTaskStatus, update requirement
      objective: Phase auto-completion works correctly with duplicate task names across plans
      changes: internal/cli/task.go (fix inner loop), requirement update
      details: Replace utils.Slugify(te.Name)==tSlug with utils.BaseSlug(te.GetBase().Slug)==tSlug and add te.PlanRef==t.PlanRef scoping
      tasks:
        - fix-phase-auto-complete-task-matching
        - update-requirement-statement
        - build-and-verify
changes:
    requirements:
        extended:
            - id: 00l7
              slug: plan-phase-auto-completes-on-final-task-done
              what: Clarify that task lookup must be scoped to the same plan
              why: Current statement doesn't specify plan scoping
              field_changes:
                statement: While a plan phase is in_progress and all of its tasks within the same plan have status completed, the syde CLI shall transition the phase to completed.
              tasks:
                - update-requirement-statement
    components:
        extended:
            - id: ze5a
              slug: cli-commands
              what: Fix setTaskStatus inner loop to scope by plan and match by slug instead of re-slugified name
              why: 'Cross-plan name collision bug: tasks from other plans with matching slugified names block auto-completion'
              tasks:
                - fix-phase-auto-complete-task-matching
                - build-and-verify
---
