---
id: TSK-0118
kind: task
name: Delete existing requirements and relax traceability
slug: delete-existing-requirements-and-relax-traceability-7gd7
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:10:16Z"
task_status: completed
priority: high
objective: All .syde/requirements/*.md files are deleted and requirementTraceFindings emits WARN (not ERROR) for task/plan/flow/design, ERROR only for component/contract/concept/system.
details: Remove .syde/requirements; edit internal/audit/graph_rules.go requirementTraceFindings to switch severity based on source entity kind.
acceptance: ls .syde/requirements is empty; syde sync check --strict shows WARNs (not ERRORs) for task/plan missing requirement links.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/graph_rules.go
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_1
created_at: "2026-04-15T09:53:21Z"
completed_at: "2026-04-15T10:10:16Z"
---
