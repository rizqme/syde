---
contract_kind: cli
description: List tasks optionally filtered by status.
id: CON-0040
input: syde task list [--status X]
input_parameters:
    - description: filter by pending|in_progress|completed|blocked|cancelled
      path: --status
      type: string
interaction_pattern: request-response
kind: contract
name: List Tasks
output: tabular task list
output_parameters:
    - description: task slug + status + priority + plan
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
slug: list-tasks-0sgj
updated_at: "2026-04-14T03:27:04Z"
---
