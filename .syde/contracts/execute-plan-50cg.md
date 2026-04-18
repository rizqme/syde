---
id: CON-0030
kind: contract
name: Execute Plan
slug: execute-plan-50cg
description: Transition an approved plan to in-progress.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde plan execute <slug>
input_parameters:
- path: slug
  type: string
  description: positional, required. Plan must be approved
output: exit 0; transitions plan_status to in-progress
output_parameters:
- path: new_status
  type: string
  description: in-progress
---
