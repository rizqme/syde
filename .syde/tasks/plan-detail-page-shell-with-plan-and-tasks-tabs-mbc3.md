---
id: TSK-0025
kind: task
name: Plan detail page shell with Plan and Tasks tabs
slug: plan-detail-page-shell-with-plan-and-tasks-tabs-mbc3
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: 'Route /__plan__/<slug> renders a plan detail page with two tabs: Plan (design + changes) and Tasks (phases with nested tasks). The active tab persists via URL query param ?tab=plan or ?tab=tasks.'
details: 'Add web/src/pages/PlanDetailScreen.tsx that loads GET /api/plan/<slug>, renders a header (name, status, progress, approved_at) and a Tabs component (reuse Radix/Headless/Tailwind pattern from existing codebase, or hand-roll a simple two-button switcher). Tab 1 embeds the existing PlanView (design + changes from task 13). Tab 2 embeds the new PhaseTaskList (task below). Breadcrumb: Plans > <plan name>.'
acceptance: Opening /__plan__/<slug> shows Plan tab by default with design and changes; clicking Tasks switches to the phases view; tab persists on reload via query param.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:44:59Z"
completed_at: "2026-04-15T12:14:54Z"
---
