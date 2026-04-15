---
id: TSK-0129
kind: task
name: Good-requirement audit rule
slug: good-requirement-audit-rule-scne
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:47:12Z"
task_status: completed
priority: high
objective: syde sync check --strict emits ERROR for any requirement whose statement is not EARS-compliant or whose req_type, priority, or verification is missing.
details: 'internal/audit/graph_rules.go: add goodRequirementFindings iterating requirement entities, reusing model.earsPatternMatch for statement check and emitting one Finding per missing field. Wire into internal/audit/audit.go Run() alongside requirementTraceFindings.'
acceptance: A requirement missing req_type is flagged; one with statement 'Add button' is flagged; one with all four fields correct passes.
affected_entities:
    - audit-engine-4ktg
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_3
created_at: "2026-04-15T09:53:58Z"
completed_at: "2026-04-15T10:47:12Z"
---
