---
contract_kind: cli
description: List all captured learnings.
id: CON-0060
input: syde learn list
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: List Learnings
output: tabular learning list
output_parameters:
    - description: learning slug + category + confidence + text
      path: rows
      type: '[]string'
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
slug: list-learnings-vuo4
updated_at: "2026-04-14T03:27:06Z"
---
