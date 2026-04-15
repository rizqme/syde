---
contract_kind: cli
description: Mark a task as in_progress.
id: CON-0037
input: syde task start <slug>
input_parameters:
    - description: positional, required
      path: slug
      type: string
interaction_pattern: request-response
kind: contract
name: Start Task
output: exit 0; transitions task_status to in_progress
output_parameters:
    - description: in_progress
      path: new_status
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: start-task-wa36
updated_at: "2026-04-14T03:27:04Z"
---
