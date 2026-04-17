---
id: CON-0032
kind: contract
name: Estimate Plan
slug: estimate-plan-kwt6
description: Check plan size and recommend splitting if too large.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde plan estimate <slug>
input_parameters:
    - path: slug
      type: string
      description: positional, required
output: Size recommendation on stdout
output_parameters:
    - path: task_count
      type: int
      description: number of tasks in the plan
    - path: recommendation
      type: string
      description: ok / consider-splitting
---
