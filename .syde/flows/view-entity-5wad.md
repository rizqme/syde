---
id: FLW-0013
kind: flow
name: View Entity
slug: view-entity-5wad
description: User views an entity's full detail
tags:
    - authoring
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: User runs syde get <slug> or clicks an entity in the dashboard
goal: User sees the entity's fields, relationships, and file refs
steps:
    - id: s1
      action: User runs syde get <slug>
      contract: get-entity
      description: CLI fetches entity via HTTP
      on_success: s2
    - id: s2
      action: System returns entity detail
      contract: get-entity-http
      description: HTTP API resolves slug and returns JSON
      on_success: done
---
