---
id: CPT-0004
kind: concept
name: Relationship
slug: relationship-hjgt
description: A typed directed link between two entities, stored on the source.
relationships:
    - target: syde
      type: belongs_to
    - target: entity
      type: relates_to
    - target: entity-model
      type: implemented_by
updated_at: "2026-04-17T08:25:52Z"
meaning: A typed directed link between two entities, stored on the source entity
invariants: Target must resolve to an existing entity at validation time. depends_on forms a DAG across components. belongs_to nesting among systems must be acyclic.
---
