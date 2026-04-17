---
id: TSK-0118
kind: task
name: Plan authoring warns on uncovered contract surfaces in new or extended requirements
slug: plan-authoring-warns-on-uncovered-contract-surfaces-in-new-or-extended-requirements-x3jm
relationships:
    - target: syde
      type: belongs_to
    - target: audit-overlap-plan-detector-coverage-symmetry-tasks
      type: references
updated_at: "2026-04-17T11:04:19Z"
task_status: completed
objective: syde plan check emits WARN when a plan's requirement lane mentions a surface not matched by any contract (new or extended) in the same diff
details: 'Extend internal/audit/plan_authoring.go planAuthoringFindings: for each new/extended requirement change, call ExtractSurfaces on its statement; build a set of contract-covered surfaces from the plan''s contract lane (new contracts'' input + input_parameters, extended contracts'' target name + any field_changes on input); emit WARN for each uncovered surface with the finding key ''requirement_contract_surface_coverage''.'
acceptance: syde plan check on a crafted plan with req mentioning 'syde foo' and no contract for 'syde foo' warns; plan with matching contract does not
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/plan_authoring.go
plan_ref: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
plan_phase: phase_9
created_at: "2026-04-17T09:46:36Z"
completed_at: "2026-04-17T10:23:22Z"
---
