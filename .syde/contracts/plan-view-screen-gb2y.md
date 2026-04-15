---
contract_kind: screen
description: Hierarchical plan detail with phases, tasks, and progress
files:
    - web/src/pages/PlanView.tsx
id: CON-0073
input: /plan/<slug>
input_parameters:
    - description: plan slug
      path: slug
      type: string
interaction_pattern: render
kind: contract
name: Plan View Screen
output: rendered plan tree
output_parameters:
    - description: navigate to task detail
      path: task-click
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
slug: plan-view-screen-gb2y
updated_at: "2026-04-15T03:08:10Z"
wireframe: '<screen name="Plan View" direction="vertical"><navbar><heading>Plan: implementation</heading><button>Approve</button></navbar><main name="Phases"><card name="Phase 1"><heading>Lexer fix</heading><list><item><label>add inTag flag</label><label>completed</label></item></list></card><card name="Phase 2"><heading>Renderer</heading><list><item><label>RenderWireframeHTML</label><label>completed</label></item><item><label>CLI command</label><label>in progress</label></item></list></card></main></screen>'
---
