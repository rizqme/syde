---
id: CON-0040
kind: contract
name: List Tasks
slug: list-tasks-0sgj
description: List tasks optionally filtered by status.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde task list [--status X]
input_parameters:
- path: --status
  type: string
  description: filter by pending|in_progress|completed|blocked|cancelled
output: tabular task list
output_parameters:
- path: rows
  type: '[]string'
  description: task slug + status + priority + plan
---
