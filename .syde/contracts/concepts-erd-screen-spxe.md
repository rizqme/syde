---
contract_kind: screen
description: React Flow canvas rendering every concept entity with attribute rows and relationship edges
files:
    - web/src/pages/ERD.tsx
id: CON-0076
input: /concept (ERD toggle active)
input_parameters:
    - description: fixed value erd
      path: view
      type: string
interaction_pattern: render
kind: contract
name: Concepts ERD Screen
output: rendered ERD canvas
output_parameters:
    - description: return to list mode and select entity
      path: node-click
      type: event
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: concepts-erd-screen-spxe
updated_at: "2026-04-15T03:08:10Z"
wireframe: <screen name="Concepts ERD" direction="vertical"><navbar><heading>Concepts</heading><button>List</button><button>ERD</button></navbar><main name="Canvas"><placeholder/></main></screen>
---
