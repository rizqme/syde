---
id: TSK-0060
kind: task
name: Create requirements for new audit rules
slug: create-requirements-for-new-audit-rules-tmi0
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-chart-and-doc-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: Two new requirements created and linked
details: syde add requirement for step-contract WARN and contract-coverage ERROR per the plan changes
acceptance: syde query --kind requirement --search 'flow step' returns both
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_2
created_at: "2026-04-16T09:22:29Z"
completed_at: "2026-04-16T10:41:08Z"
---
