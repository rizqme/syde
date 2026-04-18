---
id: CON-0012
kind: contract
name: Remove Entity
slug: remove-entity-t21l
description: Delete an entity file and its index entries.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde remove <slug> [--force]
input_parameters:
- path: slug
  type: string
  description: positional, required
- path: --force
  type: bool
  description: skip confirmation prompt
output: 'exit 0; prints ''Removed <kind>: <name>'''
output_parameters:
- path: confirmation
  type: string
  description: human confirmation line
---
