---
id: TSK-0012
kind: task
name: Expose plan with changes via HTTP API
slug: expose-plan-with-changes-via-http-api-4ke4
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: GET /api/plan/<slug> returns the full plan including design, phases, tasks, and changes — with Extended entries pre-resolved to include current entity field values for side-by-side rendering.
details: internal/dashboard/api.go add a handlePlanDetail handler that loads the plan, iterates its Changes, and for each ExtendedChange with FieldChanges resolves the target entity and attaches a CurrentValues map[string]interface{} snapshot. For NewChange the Draft is passed through as-is so the frontend can render kind-specific previews.
acceptance: curl /api/plan/<slug> returns JSON with changes.components.extended[0].current_values populated for declared field names.
affected_entities:
    - http-api-afos
affected_files:
    - internal/dashboard/api.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_4
created_at: "2026-04-15T11:41:44Z"
completed_at: "2026-04-15T12:09:23Z"
---
