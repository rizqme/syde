---
id: TSK-0083
kind: task
name: Update concept detail to glossary style
slug: update-concept-detail-to-glossary-style-dp9z
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-concept-redesign-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: Concept detail shows meaning, lifecycle, invariants, then grouped relationship chips
details: 'Update the concept case in EntityDetail KindFields. Group relationships by type: implemented_by, exposed_via, used_in, relates_to. Each group has a label and clickable chips.'
acceptance: Opening a concept shows glossary layout with grouped relationships
affected_entities:
    - web-spa-jy9z
    - entity-model-f28o
affected_files:
    - web/src/components/EntityDetail.tsx
    - internal/model/relationship.go
plan_ref: concept-entity-redesign-glossary-with-role-based-links
plan_phase: phase_3
created_at: "2026-04-16T11:18:38Z"
completed_at: "2026-04-17T08:20:45Z"
---
