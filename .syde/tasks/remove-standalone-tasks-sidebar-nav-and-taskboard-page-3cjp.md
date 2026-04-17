---
id: TSK-0023
kind: task
name: Remove standalone Tasks sidebar nav and TaskBoard page
slug: remove-standalone-tasks-sidebar-nav-and-taskboard-page-3cjp
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: The dashboard no longer shows Tasks as a standalone top-level nav — tasks live only inside the plan detail page.
details: Remove the Tasks item from web/src/components/Sidebar.tsx, delete or stub web/src/pages/TaskBoard.tsx, drop the /__task__ route from App.tsx, clean up KIND_ICON_MAP/KindBadge/EntityFilterBar entries that treat task as a separate top-level kind (tasks still render inside the plan detail page but do not need their own inbox page).
acceptance: Opening the dashboard shows no Tasks nav item; navigating to the old /__task__ path redirects to /__plan__ or 404.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:44:58Z"
completed_at: "2026-04-15T12:14:54Z"
---
