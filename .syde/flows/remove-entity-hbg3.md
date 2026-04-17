---
id: FLW-0015
kind: flow
name: Remove Entity
slug: remove-entity-hbg3
description: User deletes an entity from the design model
tags:
    - authoring
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: User runs syde remove <slug>
goal: Entity file and all index entries are removed
steps:
    - id: s1
      action: User runs syde remove
      contract: remove-entity
      description: CLI confirms and deletes via HTTP
      on_success: s2
    - id: s2
      action: System removes index entries
      description: Cleans up BadgerDB keys
      on_success: done
---
