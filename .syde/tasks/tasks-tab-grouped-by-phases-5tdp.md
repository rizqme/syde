---
id: TSK-0026
kind: task
name: Tasks tab grouped by phases
slug: tasks-tab-grouped-by-phases-5tdp
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: The Tasks tab on the plan detail page renders each phase as a collapsible section containing its nested tasks, with status icon and priority for each task.
details: 'Add web/src/components/PhaseTaskList.tsx. For each phase in plan.phases: collapsible header showing phase name + status icon + progress (N/M tasks done). Inside: each task row shows status icon (○●✓✗–), task name, priority badge, and a short objective snippet. Clicking a task opens the existing EntityDetail panel via the dashboard''s entity-detail-in-sidebar pattern. Data comes from GET /api/plan/<slug> which already exposes phases[].tasks[] — extend the handler if needed to include task fields (status, priority, objective).'
acceptance: The Tasks tab shows every phase with its tasks grouped underneath; task status icons match syde task list output; clicking a task opens its detail view.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:44:59Z"
completed_at: "2026-04-15T12:14:54Z"
---
