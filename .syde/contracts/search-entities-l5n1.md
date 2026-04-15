---
contract_kind: cli
description: Full-text search across all entities.
id: CON-0013
input: syde search <query>
input_parameters:
    - description: positional, required. Full-text query string
      path: query
      type: string
interaction_pattern: request-response
kind: contract
name: Search Entities
output: ranked list on stdout
output_parameters:
    - description: matching entity slugs with score and snippet
      path: hits
      type: '[]string'
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
slug: search-entities-l5n1
updated_at: "2026-04-14T03:27:04Z"
---
