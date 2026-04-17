---
id: TSK-0016
kind: task
name: Kind-specific New entity draft previews
slug: kind-specific-new-entity-draft-previews-tif6
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-view-shall-render-screen-wireframes-inline-mhyy
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: medium
objective: 'Non-contract NewChange drafts render kind-appropriate previews: component responsibility+capabilities+boundaries, requirement statement+type+priority+verification, concept meaning+invariants+attributes+actions, flow trigger+goal+happy_path, system purpose+quality_goals.'
details: web/src/components/NewEntityDraftView.tsx switches on the Change's target kind. Reuse styling from the existing EntityDetail case blocks for each kind where possible. Fall back to a generic key/value list for unknown kinds.
acceptance: A plan with one NewChange per non-contract kind renders previews without falling back to the generic view.
affected_entities:
    - web-spa-jy9z
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:41:44Z"
completed_at: "2026-04-15T12:19:14Z"
---
