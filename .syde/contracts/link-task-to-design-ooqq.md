---
contract_kind: cli
description: Link a task to a design entity.
id: CON-0042
input: syde task link <task-slug> <entity-slug>
input_parameters:
    - description: positional, required
      path: task-slug
      type: string
    - description: positional, required. Target design entity
      path: entity-slug
      type: string
interaction_pattern: request-response
kind: contract
name: Link Task To Design
output: exit 0; prints confirmation
output_parameters:
    - description: echoed link
      path: confirmation
      type: string
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
slug: link-task-to-design-ooqq
updated_at: "2026-04-14T03:27:05Z"
---
