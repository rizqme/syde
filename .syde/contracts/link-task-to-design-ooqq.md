---
id: CON-0042
kind: contract
name: Link Task To Design
slug: link-task-to-design-ooqq
description: Link a task to a design entity.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde task link <task-slug> <entity-slug>
input_parameters:
    - path: task-slug
      type: string
      description: positional, required
    - path: entity-slug
      type: string
      description: positional, required. Target design entity
output: exit 0; prints confirmation
output_parameters:
    - path: confirmation
      type: string
      description: echoed link
---
