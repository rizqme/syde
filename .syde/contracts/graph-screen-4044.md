---
id: CON-0072
kind: contract
name: Graph Screen
slug: graph-screen-4044
description: Concept relationship graph visualization.
files:
    - web/src/pages/Graph.tsx
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
updated_at: "2026-04-18T09:28:27Z"
contract_kind: screen
interaction_pattern: render
input: /__graph__
input_parameters:
    - path: depth
      type: int
      description: optional traversal depth, default 2
output: rendered graph canvas
output_parameters:
    - path: node-click
      type: event
      description: navigate to entity detail
wireframe: <screen name="Graph" direction="vertical"><navbar><heading>Relationship Graph</heading><button>Reset</button><button>Fit</button></navbar><main name="Canvas"><placeholder/></main></screen>
---
