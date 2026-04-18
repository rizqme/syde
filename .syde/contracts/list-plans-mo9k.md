---
id: CON-0033
kind: contract
name: List Plans
slug: list-plans-mo9k
description: List all plans with their status and progress.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde plan list
input_parameters:
- path: _
  type: '-'
  description: no arguments
output: tabular plan list with progress
output_parameters:
- path: rows
  type: '[]string'
  description: plan slug + status + completed/total tasks
---
