---
id: CON-0073
kind: contract
name: Plan View Screen
slug: plan-view-screen-gb2y
description: Hierarchical plan detail with phases, tasks, and progress; renders Design + Changes via PlanDetailScreen.
files:
    - web/src/components/PlanDetailPanel.tsx
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: web-spa
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: screen
interaction_pattern: render
input: /plan/<slug>
input_parameters:
    - path: slug
      type: string
      description: plan slug
output: rendered plan tree
output_parameters:
    - path: task-click
      type: event
      description: navigate to task detail
wireframe: '<screen name="Plans Inbox" direction="horizontal"><sidebar><list name="plans"><item><label>Plans Inbox 2-Column Layout</label><label>draft</label></item><item><label>Revamp Planning</label><label>completed</label></item></list></sidebar><main name="Plan Detail"><navbar><heading>Plan: Plans Inbox 2-Column Layout</heading><button>Approve</button></navbar><tabs><tab name="Plan"/><tab name="Tasks"/></tabs><section><heading>Design</heading><paragraph>Detailed implementation prose...</paragraph></section><section><heading>Changes</heading><tabs><tab name="Requirements"/><tab name="Components"/><tab name="Contracts"/></tabs><card><label>Extended audit-engine</label></card></section></main></screen>'
---
