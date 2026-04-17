---
id: TSK-0070
kind: task
name: Add steps to existing flows
slug: add-steps-to-existing-flows-04f0
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-flow-authoring-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: 'Verify full coverage: all 71 contracts referenced by at least one flow step'
details: Run syde sync check. If any contract is uncovered, add it to the nearest matching flow. Iterate until 0 contract-flow errors.
acceptance: syde sync check shows 0 contract-flow errors
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_4
created_at: "2026-04-16T09:23:28Z"
completed_at: "2026-04-16T10:55:59Z"
---
