---
contract_kind: cli
description: Delete an entity file and its index entries.
id: CON-0012
input: syde remove <slug> [--force]
input_parameters:
    - description: positional, required
      path: slug
      type: string
    - description: skip confirmation prompt
      path: --force
      type: bool
interaction_pattern: request-response
kind: contract
name: Remove Entity
output: 'exit 0; prints ''Removed <kind>: <name>'''
output_parameters:
    - description: human confirmation line
      path: confirmation
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: remove-entity-t21l
updated_at: "2026-04-14T03:27:04Z"
---
