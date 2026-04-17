---
id: TSK-0064
kind: task
name: Delete catch-all flow and remove contract relationships
slug: delete-catch-all-flow-and-remove-contract-relationships-64c5
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-flow-authoring-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: Design Model Operations Coverage flow is removed along with all its contract involves relationships
details: syde remove design-model-operations-coverage --force. Then update all 71 contracts to remove their involves relationship to this flow.
acceptance: syde query design-model-operations-coverage returns not found
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_4
created_at: "2026-04-16T09:23:28Z"
completed_at: "2026-04-16T10:51:25Z"
---
