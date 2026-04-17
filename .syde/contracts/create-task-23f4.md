---
id: CON-0035
kind: contract
name: Create Task
slug: create-task-23f4
description: Create a task with objective, details, acceptance, and affected entities/files.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde task create <name> [flags]
input_parameters:
    - path: name
      type: string
      description: positional, required
    - path: --plan
      type: string
      description: parent plan slug
    - path: --phase
      type: string
      description: parent plan phase ID
    - path: --priority
      type: string
      description: 'high|medium|low (default: medium)'
    - path: --objective
      type: string
      description: what this task achieves
    - path: --details
      type: string
      description: implementation specifics
    - path: --acceptance
      type: string
      description: observable done-condition
    - path: --affected-entity
      type: '[]string'
      description: repeatable slug this task will modify (validated)
    - path: --affected-file
      type: '[]string'
      description: repeatable file path this task will touch (must exist in tree)
    - path: --entity
      type: '[]string'
      description: legacy entity ref (prefer --affected-entity)
output: exit 0; prints task slug
output_parameters:
    - path: slug
      type: string
      description: new task slug
---
