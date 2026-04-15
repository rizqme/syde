---
contract_kind: cli
description: List entities optionally filtered by kind or tag.
id: CON-0010
input: syde list [kind] [--tag X]
input_parameters:
    - description: positional, optional. Filter by entity kind
      path: kind
      type: string
    - description: optional tag filter
      path: --tag
      type: string
interaction_pattern: request-response
kind: contract
name: List Entities
output: tabular list on stdout
output_parameters:
    - description: 'one row per entity: kind + name + slug'
      path: rows
      type: '[]string'
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: list-entities-0iec
updated_at: "2026-04-14T03:27:03Z"
---
