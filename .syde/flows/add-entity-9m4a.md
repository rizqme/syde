---
id: FLW-0012
kind: flow
name: Add Entity
slug: add-entity-9m4a
description: User creates a new entity in the design model
tags:
    - authoring
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T10:43:19Z"
trigger: User runs syde add <kind> <name>
goal: A new entity file exists with allocated ID, slug, and index entries
steps:
    - id: s1
      action: User runs syde add <kind> <name>
      contract: add-entity
      description: CLI validates fields and creates entity
      on_success: s2
    - id: s2
      action: System allocates counter ID
      contract: counter-key
      description: Increments per-kind counter in BadgerDB
      on_success: s3
    - id: s3
      action: System writes entity index
      contract: entity-index-key
      description: Stores FileRef under e:<kind>:<id>
      on_success: s4
    - id: s4
      action: System writes slug index
      contract: slug-index-key
      description: Maps slug to entity ID
      on_success: s5
    - id: s5
      action: System writes tag indexes
      contract: tag-index-key
      description: Indexes each tag
      on_success: s6
    - id: s6
      action: System writes word index
      contract: word-index-key
      description: Full-text indexes entity fields
      on_success: s7
    - id: s7
      action: System writes outgoing relationship index
      contract: outgoing-relationship-index-key
      description: Indexes source→target edges
      on_success: s8
    - id: s8
      action: System writes incoming relationship index
      contract: incoming-relationship-index-key
      description: Indexes target←source edges
      on_success: s9
    - id: s9
      action: For kind=requirement, CLI blocks on unacknowledged overlaps
      contract: add-requirement
      description: Runs TF-IDF check and exits non-zero unless all overlaps are acknowledged or --force
      on_success: done
---
