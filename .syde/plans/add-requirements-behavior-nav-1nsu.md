---
id: PLN-0018
kind: plan
name: Add Requirements Behavior Nav
slug: add-requirements-behavior-nav-1nsu
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: requirements-belong-in-dashboard-behavior-nav-ahyi
      type: references
    - target: approved-plan-add-requirements-behavior-nav-w4ow
      type: references
      label: requirement
updated_at: "2026-04-15T07:15:50Z"
plan_status: completed
background: Requirement entities exist in the model, but the dashboard sidebar does not expose them.
objective: Show Requirements in the Behavior section of the dashboard sidebar and route it through the generic entity list/detail flow.
scope: Update the Web SPA navigation/icon/badge styling for requirement entities; do not add a custom requirements page.
source: manual
created_at: "2026-04-15T07:14:01Z"
approved_at: "2026-04-15T07:14:21Z"
completed_at: "2026-04-15T07:15:50Z"
phases:
    - id: phase_1
      name: Sidebar navigation
      status: completed
      description: Expose requirements in the dashboard entity navigation.
      objective: Requirements appear under Behavior and use the generic entity list/detail route.
      changes: Update sidebar kind groups, requirement icon mapping, badge color, and kind accent styling.
      details: Add requirement to the Behavior group in Sidebar; add a requirement icon and iconForKind mapping; add KindBadge and CSS requirement accents.
      tasks:
        - expose-requirements-in-behavior-sidebar
---
