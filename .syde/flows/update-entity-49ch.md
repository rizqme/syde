---
id: FLW-0014
kind: flow
name: Update Entity
slug: update-entity-49ch
description: User modifies an existing entity's fields or relationships
tags:
    - authoring
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: User runs syde update <slug> with changed flags
goal: Entity file and indexes reflect the updated state
steps:
    - id: s1
      action: User runs syde update
      contract: update-entity
      description: CLI sends updated fields via HTTP
      on_success: s2
    - id: s2
      action: System re-indexes entity
      description: Updates affected BadgerDB keys
      on_success: done
---
