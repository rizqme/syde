---
id: TSK-0057
kind: task
name: Add --step CLI flag for flow add and update
slug: add-step-cli-flag-for-flow-add-and-update-01le
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-flow-authoring-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: syde add flow and syde update support --step repeatable flag
details: 'Add addFlowSteps []string flag with pipe-separated parsing. Parse into []FlowStep in add.go and update.go flow cases. Format: action|contract|description|on_success|on_failure (contract and later fields optional).'
acceptance: syde add flow Test --step 'do thing|some-contract|desc|next|fail' creates flow with one step
affected_entities:
    - cli-commands-hpjb
affected_files:
    - internal/cli/add.go
    - internal/cli/update.go
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_1
created_at: "2026-04-16T09:22:11Z"
completed_at: "2026-04-16T10:37:49Z"
---
