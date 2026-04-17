---
id: TSK-0042
kind: task
name: Update Plan View Screen contract wireframe and files list
slug: update-plan-view-screen-contract-wireframe-and-files-list-c82a
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: 'plan-view-screen-gb2y reflects the new inbox layout: files list points at PlanDetailPanel.tsx and wireframe depicts the 2-column horizontal inbox with sidebar list + main detail tabs.'
details: syde update plan-view-screen --file web/src/components/PlanDetailPanel.tsx (after removing the old PlanDetailScreen.tsx file ref), and syde update plan-view-screen --wireframe '<screen direction=horizontal>...</screen>' with the inbox layout. The wireframe value matches the field_changes declared in the plan's contract Extended entry.
acceptance: Querying plan-view-screen returns the new files list and wireframe; the Extended FieldChanges in the plan match the actual contract state.
affected_entities:
    - plan-view-screen-gb2y
affected_files:
    - web/src/components/PlanDetailPanel.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_4
created_at: "2026-04-15T13:18:13Z"
completed_at: "2026-04-15T21:35:39Z"
---
