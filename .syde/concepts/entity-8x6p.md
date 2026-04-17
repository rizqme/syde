---
id: CPT-0001
kind: concept
name: Entity
slug: entity-8x6p
description: A typed, persisted node in the syde design model.
relationships:
    - target: syde
      type: belongs_to
    - target: add-entity
      type: used_in
    - target: entity-model
      type: implemented_by
updated_at: "2026-04-17T08:25:52Z"
meaning: A typed, persisted node in the syde design model — system, component, contract, concept, flow, decision, plan, task, design, or requirement
lifecycle: draft (in memory) → created via 'syde add' with allocated ID + slug → updated via 'syde update' → removed via 'syde remove'. IDs are never reused.
invariants: ID is unique within kind. Slug filename is unique on disk. Every relationship target must resolve to an existing entity. Component must have purpose, responsibility, ≥1 capability. Contract must have input, ≥1 input_parameter, output, ≥1 output_parameter.
---
