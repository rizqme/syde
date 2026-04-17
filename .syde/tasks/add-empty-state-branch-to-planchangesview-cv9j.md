---
id: TSK-0030
kind: task
name: Add empty-state branch to PlanChangesView
slug: add-empty-state-branch-to-planchangesview-cv9j
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-changes-view-shall-group-extended-by-target-6qld
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: PlanChangesView shows a muted 'No changes declared yet' card when every lane is empty, instead of rendering nothing.
details: 'After building the lane list, if total entries === 0, render a centered card with the empty-state message and a hint: ''Use syde plan add-change to record deletions, extensions, or new entities.'' Match the dashed-border muted styling of EntityEmptyState.'
acceptance: Opening a plan with an empty changes block shows the empty state instead of a blank section.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/PlanChangesView.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_3
created_at: "2026-04-15T13:03:56Z"
completed_at: "2026-04-15T15:09:37Z"
---
