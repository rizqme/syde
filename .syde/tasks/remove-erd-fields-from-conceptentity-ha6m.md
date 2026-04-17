---
id: TSK-0078
kind: task
name: Remove ERD fields from ConceptEntity
slug: remove-erd-fields-from-conceptentity-ha6m
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-concept-redesign-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: ConceptEntity has only Meaning, Invariants, Lifecycle
details: Remove Attributes, Actions, ConceptRelationships, DataSensitivity, StructureNotes and their parse helpers
acceptance: go build clean; ConceptEntity has 3 domain fields
affected_entities:
    - entity-model-f28o
    - entity-model
affected_files:
    - internal/model/entity.go
    - internal/model/validation.go
plan_ref: concept-entity-redesign-glossary-with-role-based-links
plan_phase: phase_1
created_at: "2026-04-16T11:18:38Z"
completed_at: "2026-04-17T02:49:27Z"
---
