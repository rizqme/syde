---
id: TSK-0028
kind: task
name: Wire PlanDetailPanel into the right column for plan kind
slug: wire-plandetailpanel-into-the-right-column-for-plan-kind-u71f
relationships:
    - target: plans-inbox-2-column-layout-fud8
      type: belongs_to
    - target: plans-shall-render-via-the-canonical-2-column-inbox-63p0
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: Inside the default 2-column branch, when activeKind === 'plan' the right column renders PlanDetailPanel instead of EntityDetail.
details: 'Add a conditional inside the existing right-column block: if activeKind === ''plan'' && selectedSlug, render <PlanDetailPanel slug={selectedSlug} onClose={...} onNavigate={handleNavigateEntity} onOpenFile={handleOpenFile} />. Otherwise keep the existing EntityDetail call.'
acceptance: Selecting a plan from the left list renders PlanDetailPanel inline; selecting a component still renders EntityDetail.
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/App.tsx
    - web/src/components/PlanDetailPanel.tsx
plan_ref: plans-inbox-2-column-layout-fud8
plan_phase: phase_1
created_at: "2026-04-15T13:03:56Z"
completed_at: "2026-04-15T15:06:28Z"
---
