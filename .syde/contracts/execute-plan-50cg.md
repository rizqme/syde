---
contract_kind: cli
description: Transition an approved plan to in-progress.
id: CON-0030
input: syde plan execute <slug>
input_parameters:
    - description: positional, required. Plan must be approved
      path: slug
      type: string
interaction_pattern: request-response
kind: contract
name: Execute Plan
output: exit 0; transitions plan_status to in-progress
output_parameters:
    - description: in-progress
      path: new_status
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: execute-plan-50cg
updated_at: "2026-04-14T03:27:04Z"
---
