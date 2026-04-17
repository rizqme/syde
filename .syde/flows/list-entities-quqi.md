---
id: FLW-0016
kind: flow
name: List Entities
slug: list-entities-quqi
description: User lists all entities of a given kind
tags:
    - authoring
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User runs syde list <kind> or views a kind page in the dashboard
goal: User sees all entities of the requested kind
steps:
    - id: s1
      action: User runs syde list <kind>
      contract: list-entities
      description: CLI fetches list via HTTP
      on_success: s2
    - id: s2
      action: System returns entity list
      contract: list-entities-http
      description: HTTP API filters by kind and returns summaries
      on_success: done
---
