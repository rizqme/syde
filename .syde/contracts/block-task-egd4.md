---
contract_kind: cli
description: Mark a task as blocked with an optional reason.
id: CON-0039
input: syde task block <slug> [--reason ...]
input_parameters:
    - description: positional, required
      path: slug
      type: string
    - description: optional block reason stored in task notes
      path: --reason
      type: string
interaction_pattern: request-response
kind: contract
name: Block Task
output: exit 0; transitions task_status to blocked
output_parameters:
    - description: blocked
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
slug: block-task-egd4
updated_at: "2026-04-14T03:27:04Z"
---
