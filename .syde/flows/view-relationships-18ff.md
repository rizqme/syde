---
id: FLW-0018
kind: flow
name: View Relationships
slug: view-relationships-18ff
description: User explores entity relationships as a graph
tags:
    - querying
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User runs syde graph <slug> or opens the graph page in the dashboard
goal: User sees the entity's relationship graph with neighbors
steps:
    - id: s1
      action: User runs syde graph <slug>
      contract: relationship-graph
      description: CLI renders relationship graph
      on_success: s2
    - id: s2
      action: Dashboard renders graph view
      contract: graph-screen
      description: Interactive node-edge graph
      on_success: done
---
