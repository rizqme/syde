---
id: TSK-0009
kind: task
name: Implement planCompletionFindings audit rule
slug: implement-plancompletionfindings-audit-rule-20dn
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plan-changes-shall-list-their-implementing-tasks-gb2a
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: Audit produces ERROR findings for every drift between a plan's declared Changes and actual entity state.
details: 'New file internal/audit/plan_completion.go with planCompletionFindings function. Walk every plan whose status is approved or completed; for each change:\n- DeletedChange: if target slug resolves in the audit graph, emit ERROR ''plan claims deletion but entity still exists''.\n- NewChange: if no entity of the declared kind with a matching name (case-insensitive slugified) exists, emit ERROR ''plan claims new entity but it was not created''.\n- ExtendedChange with FieldChanges: resolve target entity, compare declared fields against current values via reflection on the BaseEntity + kind-specific struct, emit ERROR per mismatching field.\n- ExtendedChange without FieldChanges: emit WARN ''no field-level diff declared — hand review required''.'
acceptance: 'Unit-equivalent: a plan with a DeletedChange targeting an existing entity triggers ERROR in syde sync check --strict.'
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/plan_completion.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_3
created_at: "2026-04-15T11:41:13Z"
completed_at: "2026-04-15T12:00:00Z"
---
