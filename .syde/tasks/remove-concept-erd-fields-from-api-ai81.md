---
id: TSK-0081
kind: task
name: Remove concept ERD fields from API
slug: remove-concept-erd-fields-from-api-ai81
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-concept-redesign-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: EntitySummary and FormatJSON no longer include attributes/actions
details: Remove attributes/actions from query engine Filter and formatter FormatJSON
acceptance: curl entity/<concept> returns no attributes/actions keys
affected_entities:
    - query-engine-9k84
affected_files:
    - internal/query/engine.go
    - internal/query/formatter.go
plan_ref: concept-entity-redesign-glossary-with-role-based-links
plan_phase: phase_2
created_at: "2026-04-16T11:18:38Z"
completed_at: "2026-04-17T08:13:52Z"
---
