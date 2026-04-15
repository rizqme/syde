---
id: TSK-0107
kind: task
name: Expose Requirements in Behavior sidebar
slug: expose-requirements-in-behavior-sidebar-03vy
relationships:
    - target: add-requirements-behavior-nav-1nsu
      type: belongs_to
    - target: requirements-belong-in-dashboard-behavior-nav-ahyi
      type: references
updated_at: "2026-04-15T07:15:47Z"
task_status: completed
priority: high
objective: Add Requirements to the Behavior sidebar group with icon, badge, and accent support.
details: Edit Sidebar, icons, KindBadge, and CSS theme tokens so requirement entities appear and render consistently through the generic entity browser.
acceptance: Sidebar shows Requirements under Behavior with its entity count; requirement list/detail views render with requirement icon/badge styling.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/Sidebar.tsx
    - web/src/components/icons.tsx
    - web/src/components/KindBadge.tsx
    - web/src/components/EntityEmptyState.tsx
    - web/src/index.css
plan_ref: add-requirements-behavior-nav-1nsu
plan_phase: phase_1
created_at: "2026-04-15T07:14:16Z"
completed_at: "2026-04-15T07:15:47Z"
---
