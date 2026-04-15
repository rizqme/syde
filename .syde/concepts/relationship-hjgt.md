---
attributes:
    - description: slug of the entity declaring the relationship
      name: source
      refs:
        - entity
    - description: belongs_to
      name: type
      refs:
        - depends_on|references|relates_to|exposes|applies_to|involves
    - description: slug or ID of the target entity
      name: target
      refs:
        - entity
    - description: optional free-text label (e.g. cardinality one-to-many on relates_to)
      name: label
description: A typed directed link between two entities, stored on the source.
id: CPT-0004
invariants: Target must resolve to an existing entity at validation time. depends_on forms a DAG across components. belongs_to nesting among systems must be acyclic.
kind: concept
meaning: A typed directed link between two entities, stored on the source entity
name: Relationship
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: references
    - target: entity
      type: relates_to
slug: relationship-hjgt
structure_notes: Pair of (target, type). Target is a slug or ID. Type is one of belongs_to, depends_on, exposes, consumes, uses, involves, references, relates_to, implements, applies_to, modifies, visualizes.
updated_at: "2026-04-14T10:48:03Z"
---
