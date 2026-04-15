---
category: pattern
confidence: medium
description: 'Naive grid placement in React Flow looks terrible once you have more than ~8 nodes with cross edges — lines criss-cross through node centers. Fix: run a layered graph layout (dagre is ~40KB, zero-config for LR/TB), map its output to node positions, set targetPosition/sourcePosition to Left/Right (for LR) or Top/Bottom (for TB), and switch all edges to type=''smoothstep'' for right-angle ERD-style routing. dagre''s nodesep + ranksep params control the spacing between siblings and between layers respectively.'
discovered_at: "2026-04-14T11:02:08Z"
entity_refs:
    - web-spa
id: LRN-0022
kind: learning
name: Naive grid placement in React Flow looks terrible once you h
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: naive-grid-placement-in-react-flow-looks-terrible-once-you-h-ozac
source: session-observation
updated_at: "2026-04-14T11:02:08Z"
---
