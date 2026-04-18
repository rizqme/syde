---
id: CON-0041
kind: contract
name: Create Subtask
slug: create-subtask-0fj9
description: Create a child task under a parent task.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde task sub <parent-slug> <name>
input_parameters:
- path: parent-slug
  type: string
  description: positional, required
- path: name
  type: string
  description: positional, required
output: exit 0; prints new subtask slug
output_parameters:
- path: slug
  type: string
  description: new subtask slug
---
