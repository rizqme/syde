---
id: CON-0039
kind: contract
name: Block Task
slug: block-task-egd4
description: Mark a task as blocked with an optional reason.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde task block <slug> [--reason ...]
input_parameters:
    - path: slug
      type: string
      description: positional, required
    - path: --reason
      type: string
      description: optional block reason stored in task notes
output: exit 0; transitions task_status to blocked
output_parameters:
    - path: new_status
      type: string
      description: blocked
---
