---
contract_kind: cli
description: Create a child task under a parent task.
id: CON-0041
input: syde task sub <parent-slug> <name>
input_parameters:
    - description: positional, required
      path: parent-slug
      type: string
    - description: positional, required
      path: name
      type: string
interaction_pattern: request-response
kind: contract
name: Create Subtask
output: exit 0; prints new subtask slug
output_parameters:
    - description: new subtask slug
      path: slug
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: create-subtask-0fj9
updated_at: "2026-04-14T03:27:04Z"
---
