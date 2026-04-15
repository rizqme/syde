---
contract_kind: cli
description: List all plans with their status and progress.
id: CON-0033
input: syde plan list
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: List Plans
output: tabular plan list with progress
output_parameters:
    - description: plan slug + status + completed/total tasks
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
slug: list-plans-mo9k
updated_at: "2026-04-14T03:27:04Z"
---
