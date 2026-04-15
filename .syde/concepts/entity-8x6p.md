---
attributes:
    - description: counter-based ID, globally unique within kind, never reused
      name: id
    - description: one of system/component/contract/concept/flow/decision/plan/task/design/requirement
      name: kind
    - description: filename slug with 4-char random suffix, unique on disk
      name: slug
    - description: human-readable name
      name: name
    - description: one-sentence elevator pitch, required on every entity
      name: description
    - description: concrete source file paths this entity maps to
      name: files
    - description: typed directed links to other entities
      name: relationships
      refs:
        - relationship
    - description: last-write wall clock, bumped on every save
      name: updated_at
data_sensitivity: Non-sensitive — intended for git commits. Never store secrets in entity bodies.
description: A typed, persisted node in the syde design model.
id: CPT-0001
invariants: ID is unique within kind. Slug filename is unique on disk. Every relationship target must resolve to an existing entity. Component must have purpose, responsibility, ≥1 capability. Contract must have input, ≥1 input_parameter, output, ≥1 output_parameter.
kind: concept
lifecycle: draft (in memory) → created via 'syde add' with allocated ID + slug → updated via 'syde update' → removed via 'syde remove'. IDs are never reused.
meaning: A typed, persisted node in the syde design model — system, component, contract, concept, flow, decision, plan, task, design, or requirement
name: Entity
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: entity-8x6p
structure_notes: Every entity is a markdown file with YAML frontmatter + free-form body. Shared BaseEntity fields (id, kind, name, slug, description, purpose, tags, files, notes, relationships, updated_at) are embedded by every kind-specific struct.
updated_at: "2026-04-14T10:48:02Z"
---
