---
id: TSK-0065
kind: task
name: Author Entity CRUD Lifecycle flow with steps
slug: author-entity-crud-lifecycle-flow-with-steps-bfgp
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-flow-authoring-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: 'Author planning flows (6): Create Plan, Review Plan, Approve Plan, Execute Plan, Sync Plan from Claude, plus steps on existing plan-lifecycle'
details: 'Batch script. Each flow: trigger, goal, tags=[planning], steps referencing contracts. Covers: create-plan, add-plan-phase, create-task, create-subtask, estimate-plan, show-plan, plan-view-screen, list-plans, approve-plan, execute-plan, start-task, complete-task, block-task, update-task, link-task-to-design, list-tasks, sync-plan-from-claude, update-plan, update-plan-phase'
acceptance: All planning contracts appear in flow steps
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_4
created_at: "2026-04-16T09:23:28Z"
completed_at: "2026-04-16T10:52:21Z"
---
