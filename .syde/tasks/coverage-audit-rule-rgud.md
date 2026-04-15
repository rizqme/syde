---
id: TSK-0130
kind: task
name: Coverage audit rule
slug: coverage-audit-rule-rgud
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:47:59Z"
task_status: completed
priority: high
objective: syde sync check --strict emits ERROR for any component, contract, concept, or system entity not connected to at least one requirement.
details: 'internal/audit/graph_rules.go: add coverageFindings that builds a set of (entity -> requirement) edges in either direction (outbound from entity to requirement, or inbound from requirement via refines) and reports any component/contract/concept/system missing from the set. Wire into audit.go Run(). Flows/plans/tasks/designs excluded per user scope.'
acceptance: A bare component with no requirement links is flagged; covering it with a single requirement clears the finding.
affected_entities:
    - audit-engine-4ktg
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_3
created_at: "2026-04-15T09:53:58Z"
completed_at: "2026-04-15T10:47:59Z"
---
