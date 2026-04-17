---
id: TSK-0031
kind: task
name: Delete dead PlansInboxScreen and PlanDetailScreen files
slug: delete-dead-plansinboxscreen-and-plandetailscreen-files-z8g2
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: Both dead page files are deleted and nothing imports them anymore.
details: rm web/src/pages/PlansInboxScreen.tsx web/src/pages/PlanDetailScreen.tsx. Remove their imports from App.tsx. Run rg PlansInboxScreen|PlanDetailScreen web/src to confirm no matches.
acceptance: Both files gone; rg returns zero matches; bun run build clean.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/App.tsx
    - web/src/components/PlanDetailPanel.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_4
created_at: "2026-04-15T13:03:56Z"
completed_at: "2026-04-15T21:35:46Z"
---
