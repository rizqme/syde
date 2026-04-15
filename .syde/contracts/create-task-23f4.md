---
contract_kind: cli
description: Create a task with objective, details, acceptance, and affected entities/files.
id: CON-0035
input: syde task create <name> [flags]
input_parameters:
    - description: positional, required
      path: name
      type: string
    - description: parent plan slug
      path: --plan
      type: string
    - description: parent plan phase ID
      path: --phase
      type: string
    - description: 'high|medium|low (default: medium)'
      path: --priority
      type: string
    - description: what this task achieves
      path: --objective
      type: string
    - description: implementation specifics
      path: --details
      type: string
    - description: observable done-condition
      path: --acceptance
      type: string
    - description: repeatable slug this task will modify (validated)
      path: --affected-entity
      type: '[]string'
    - description: repeatable file path this task will touch (must exist in tree)
      path: --affected-file
      type: '[]string'
    - description: legacy entity ref (prefer --affected-entity)
      path: --entity
      type: '[]string'
interaction_pattern: request-response
kind: contract
name: Create Task
output: exit 0; prints task slug
output_parameters:
    - description: new task slug
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
slug: create-task-23f4
updated_at: "2026-04-14T03:27:04Z"
---
