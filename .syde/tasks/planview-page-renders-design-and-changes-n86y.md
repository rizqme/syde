---
id: TSK-0013
kind: task
name: PlanView page renders design and changes
slug: planview-page-renders-design-and-changes-n86y
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: web/src/pages/PlanView.tsx shows a Design section (markdown), existing phases/tasks, and a new Changes section grouped by kind.
details: Factor Changes rendering into a new PlanChanges component. Group by kind with tabs (Requirements / Systems / Concepts / Components / Contracts / Flows). Each lane shows Deleted/Extended/New as collapsible cards with what and why text.
acceptance: Opening /plans/<slug> in the dashboard shows Design prose + kind-tabbed change lists when changes are present.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:41:44Z"
completed_at: "2026-04-15T12:19:14Z"
---
