---
id: CON-0036
kind: contract
name: Update Task
slug: update-task-6y9m
description: Edit a task's fields or replace its affected entities/files lists.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde task update <slug> [flags]
input_parameters:
- path: slug
  type: string
  description: positional, required
- path: --objective
  type: string
  description: update objective
- path: --details
  type: string
  description: update details
- path: --acceptance
  type: string
  description: update acceptance criteria
- path: --priority
  type: string
  description: update priority
- path: --description
  type: string
  description: update description
- path: --affected-entity
  type: '[]string'
  description: replace affected entities list
- path: --affected-file
  type: '[]string'
  description: replace affected files list
output: exit 0; prints updated slug
output_parameters:
- path: slug
  type: string
  description: echoed slug
---
