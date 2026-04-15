---
id: TSK-0108
kind: task
name: Require outbound requirement links in audit
slug: require-outbound-requirement-links-in-audit-i095
relationships:
    - target: enforce-outgoing-requirement-traceability-lo0c
      type: belongs_to
    - target: strict-sync-requires-outgoing-requirement-traceability-yvdj
      type: references
updated_at: "2026-04-15T08:08:48Z"
task_status: completed
priority: high
objective: Make requirement traceability count only outgoing relationships to requirement entities.
details: Update internal/audit/graph_rules.go so each non-requirement entity must have a relationship whose target resolves to a requirement.
acceptance: An entity mentioned only by an inbound requirement relationship still fails requirement traceability.
affected_entities:
    - audit-engine-4ktg
affected_files:
    - internal/audit/graph_rules.go
plan_ref: enforce-outgoing-requirement-traceability-lo0c
plan_phase: phase_1
created_at: "2026-04-15T08:08:13Z"
completed_at: "2026-04-15T08:08:48Z"
---
