---
contract_kind: screen
description: Timeline of captured learnings sorted by recency
files:
    - web/src/pages/LearningFeed.tsx
id: CON-0075
input: /learning
input_parameters:
    - description: optional category filter
      path: category
      type: string
interaction_pattern: render
kind: contract
name: Learning Feed Screen
output: rendered learning feed
output_parameters:
    - description: open learning detail
      path: learning-click
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
slug: learning-feed-screen-p5gn
updated_at: "2026-04-15T03:08:10Z"
wireframe: <screen name="Learning Feed" direction="vertical"><navbar><heading>Learnings</heading><button>New</button></navbar><main name="Feed"><list><item><image/><label>UIML lexer attribute bug</label><label>gotcha</label></item><item><image/><label>React Flow drag fix</label><label>pattern</label></item><item><image/><label>Storage engine cascade</label><label>insight</label></item></list></main></screen>
---
