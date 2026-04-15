---
id: TSK-0114
kind: task
name: Remove learning dashboard UI
slug: remove-learning-dashboard-ui-oz5o
relationships:
    - target: remove-learning-entity-altogether-wk7c
      type: references
    - target: limit-baseline-style-requirement-fanout-m156
      type: references
    - target: remove-learning-entity-and-cap-requirement-fanout-wqyz
      type: belongs_to
updated_at: "2026-04-15T09:23:33Z"
task_status: completed
priority: high
objective: Remove learning navigation, route, page, API client types, icons, and styling from the Web SPA.
details: Update App, Sidebar, icons, API types, EntityDetail palettes, empty state accents, badges, CSS, and delete LearningFeed.
acceptance: rg finds no LearningFeed, learning nav item, or kind-learning UI references.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/App.tsx
    - web/src/components/Sidebar.tsx
    - web/src/components/icons.tsx
    - web/src/components/KindBadge.tsx
    - web/src/components/EntityEmptyState.tsx
    - web/src/components/EntityDetail.tsx
    - web/src/lib/api.ts
    - web/src/index.css
    - web/src/pages/LearningFeed.tsx
plan_ref: remove-learning-entity-and-cap-requirement-fanout-wqyz
plan_phase: phase_2
created_at: "2026-04-15T08:23:05Z"
completed_at: "2026-04-15T09:08:16Z"
---
