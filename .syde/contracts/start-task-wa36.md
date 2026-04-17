---
id: CON-0037
kind: contract
name: Start Task
slug: start-task-wa36
description: Mark a task as in_progress.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde task start <slug>
input_parameters:
    - path: slug
      type: string
      description: positional, required
output: exit 0; transitions task_status to in_progress
output_parameters:
    - path: new_status
      type: string
      description: in_progress
---
