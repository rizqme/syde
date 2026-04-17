---
id: TSK-0027
kind: task
name: Drop plan from SPECIAL_VIEWS in App.tsx
slug: drop-plan-from-specialviews-in-apptsx-rujs
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: App.tsx no longer treats plan as a special view; activeKind=plan flows through the default 2-column branch.
details: Remove 'plan' from the SPECIAL_VIEWS array. Verify the floating EntityDetail at the bottom of App.tsx no longer fires for plan kind.
acceptance: rg SPECIAL_VIEWS web/src returns no plan reference; bun run build clean.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/App.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_1
created_at: "2026-04-15T13:03:56Z"
completed_at: "2026-04-15T15:06:00Z"
---
