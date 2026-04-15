---
contract_kind: cli
description: Check plan size and recommend splitting if too large.
id: CON-0032
input: syde plan estimate <slug>
input_parameters:
    - description: positional, required
      path: slug
      type: string
interaction_pattern: request-response
kind: contract
name: Estimate Plan
output: Size recommendation on stdout
output_parameters:
    - description: number of tasks in the plan
      path: task_count
      type: int
    - description: ok / consider-splitting
      path: recommendation
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: estimate-plan-kwt6
updated_at: "2026-04-14T03:27:04Z"
---
