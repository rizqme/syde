---
id: TSK-0062
kind: task
name: Integrate FlowChart into flow detail panel
slug: integrate-flowchart-into-flow-detail-panel-5vt1
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-chart-and-doc-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: Flow entities show the flowchart in their detail panel
details: In EntityDetail or a flow-specific detail component, detect kind=flow and render FlowChart when steps exist. Keep the existing prose fields (trigger, goal, narrative) above the chart.
acceptance: Clicking a flow in the Flows inbox shows prose fields + flowchart
affected_entities:
    - web-spa-jy9z
affected_files:
    - web/src/components/EntityDetail.tsx
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_3
created_at: "2026-04-16T09:22:48Z"
completed_at: "2026-04-16T10:48:41Z"
---
