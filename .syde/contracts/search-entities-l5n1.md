---
id: CON-0013
kind: contract
name: Search Entities
slug: search-entities-l5n1
description: Full-text search across all entities.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde search <query>
input_parameters:
- path: query
  type: string
  description: positional, required. Full-text query string
output: ranked list on stdout
output_parameters:
- path: hits
  type: '[]string'
  description: matching entity slugs with score and snippet
---
