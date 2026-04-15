---
contract_kind: cli
description: Edit a task's fields or replace its affected entities/files lists.
id: CON-0036
input: syde task update <slug> [flags]
input_parameters:
    - description: positional, required
      path: slug
      type: string
    - description: update objective
      path: --objective
      type: string
    - description: update details
      path: --details
      type: string
    - description: update acceptance criteria
      path: --acceptance
      type: string
    - description: update priority
      path: --priority
      type: string
    - description: update description
      path: --description
      type: string
    - description: replace affected entities list
      path: --affected-entity
      type: '[]string'
    - description: replace affected files list
      path: --affected-file
      type: '[]string'
interaction_pattern: request-response
kind: contract
name: Update Task
output: exit 0; prints updated slug
output_parameters:
    - description: echoed slug
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
slug: update-task-6y9m
updated_at: "2026-04-14T03:27:04Z"
---
