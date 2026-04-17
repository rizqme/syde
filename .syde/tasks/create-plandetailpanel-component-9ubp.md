---
id: TSK-0029
kind: task
name: Create PlanDetailPanel component
slug: create-plandetailpanel-component-9ubp
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: web/src/components/PlanDetailPanel.tsx exists and renders the same content as PlanDetailScreen but as a prop-driven inline component.
details: |-
    Copy PlanDetailScreen body into a new component file under web/src/components/PlanDetailPanel.tsx. Replace useParams with a slug prop. Replace useNavigate with onNavigate / onClose / onOpenFile callbacks. Keep the ?tab= query-param tab persistence — reading window.location.search and setting it via history.replaceState is fine, no router needed. Header: name, status pill, progress bar, approved_at.

    Section order in the Plan tab body MUST be: (1) Background, (2) Objective, (3) Scope, (4) Design, (5) Changes. Design comes AFTER background/objective/scope so reviewers read context first and dive into the implementation prose only once they understand what the plan is for. The current PlanDetailScreen has Design first — fix that ordering when porting to PlanDetailPanel.
acceptance: Component compiles, renders the Plan tab (Design + Changes via PlanChangesView) and Tasks tab (PhaseTaskList) correctly when given a plan slug.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/PlanDetailPanel.tsx
    - web/src/App.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_2
created_at: "2026-04-15T13:03:56Z"
completed_at: "2026-04-15T15:06:28Z"
---
