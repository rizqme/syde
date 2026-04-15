---
contract_kind: screen
description: d3 force-directed visualization of every entity relationship in the project
files:
    - web/src/pages/Graph.tsx
id: CON-0072
input: /__graph__
input_parameters:
    - description: optional traversal depth, default 2
      path: depth
      type: int
interaction_pattern: render
kind: contract
name: Graph Screen
output: rendered graph canvas
output_parameters:
    - description: navigate to entity detail
      path: node-click
      type: event
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: graph-screen-4044
updated_at: "2026-04-15T03:08:10Z"
wireframe: <screen name="Graph" direction="vertical"><navbar><heading>Relationship Graph</heading><button>Reset</button><button>Fit</button></navbar><main name="Canvas"><placeholder/></main></screen>
---
