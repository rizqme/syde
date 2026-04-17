---
id: TSK-0047
kind: task
name: Group Extended changes by target slug in PlanChangesView
slug: group-extended-changes-by-target-slug-in-planchangesview-h92q
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plan-changes-view-shall-group-extended-by-target-6qld
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: PlanChangesView renders one card per (kind, target slug) for Extended changes, even when multiple Extended entries target the same entity.
details: 'web/src/components/PlanChangesView.tsx: before rendering Extended[], group by slug. Each card shows the slug header once and stacks each entry''s what/why/field_changes inside as labeled sub-sections (one section per change ID, so the user can still see they were authored as separate entries). Field changes are merged into a single field-diff table per card. The data model is unchanged — grouping happens at render time only.'
acceptance: A plan with two Extended changes on the same contract slug renders as one card with both what/why blocks visible inside.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/PlanChangesView.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:29:36Z"
completed_at: "2026-04-15T15:09:37Z"
---
