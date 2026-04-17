---
id: TSK-0080
kind: task
name: Simplify concept audit rules
slug: simplify-concept-audit-rules-azxm
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-concept-redesign-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: Audit requires meaning, recommends invariants; no attribute/cardinality checks
details: Rewrite conceptFindings in concepts.go. Remove cardinality validation from graph_rules.go.
acceptance: syde sync check passes with concepts that have no attributes
affected_entities:
    - audit-engine-4ktg
    - audit-engine
affected_files:
    - internal/audit/concepts.go
    - internal/audit/audit.go
plan_ref: concept-entity-redesign-glossary-with-role-based-links
plan_phase: phase_2
created_at: "2026-04-16T11:18:38Z"
completed_at: "2026-04-17T02:53:50Z"
---
