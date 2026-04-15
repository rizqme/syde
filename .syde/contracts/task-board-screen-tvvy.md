---
contract_kind: screen
description: Kanban-style task board grouped by status
files:
    - web/src/pages/TaskBoard.tsx
id: CON-0074
input: /task
input_parameters:
    - description: optional status filter
      path: filter
      type: string
interaction_pattern: render
kind: contract
name: Task Board Screen
output: rendered task board
output_parameters:
    - description: open task detail
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
slug: task-board-screen-tvvy
updated_at: "2026-04-15T03:08:10Z"
wireframe: <screen name="Task Board" direction="vertical"><navbar><heading>Tasks</heading><button>New</button></navbar><main name="Columns"><layout direction="horizontal"><panel name="Pending"><heading>Pending</heading><list><item><label>task one</label></item><item><label>task two</label></item></list></panel><panel name="In Progress"><heading>In Progress</heading><list><item><label>task three</label></item></list></panel><panel name="Done"><heading>Done</heading><list><item><label>task four</label></item><item><label>task five</label></item></list></panel></layout></main></screen>
---
