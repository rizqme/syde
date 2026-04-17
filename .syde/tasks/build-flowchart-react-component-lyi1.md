---
id: TSK-0063
kind: task
name: Build FlowChart React component
slug: build-flowchart-react-component-lyi1
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-chart-and-doc-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: FlowChart component renders steps as connected nodes with contract chips and directed edges
details: 'Create web/src/components/FlowChart.tsx. Each step renders as a card node: action as title, contract as clickable chip, description as body. Edges connect steps via on_success (solid green arrow) and on_failure (dashed red arrow). Layout: vertical flow with branching. Use SVG or CSS for edges.'
acceptance: Opening a flow with steps in the dashboard shows the flowchart
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/FlowChart.tsx
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_3
created_at: "2026-04-16T09:22:59Z"
completed_at: "2026-04-16T10:48:49Z"
---
