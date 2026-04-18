---
id: CON-0010
kind: contract
name: List Entities
slug: list-entities-0iec
description: List entities optionally filtered by kind or tag.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:15Z'
contract_kind: cli
interaction_pattern: request-response
input: syde list [kind] [--tag X]
input_parameters:
- path: kind
  type: string
  description: positional, optional. Filter by entity kind
- path: --tag
  type: string
  description: optional tag filter
output: tabular list on stdout
output_parameters:
- path: rows
  type: '[]string'
  description: 'one row per entity: kind + name + slug'
---
