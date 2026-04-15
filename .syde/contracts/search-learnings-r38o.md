---
contract_kind: cli
description: Search learning text by query string.
id: CON-0062
input: syde learn search <text>
input_parameters:
    - description: positional, required
      path: text
      type: string
interaction_pattern: request-response
kind: contract
name: Search Learnings
output: matching learnings on stdout
output_parameters:
    - description: matches
      path: rows
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
slug: search-learnings-r38o
updated_at: "2026-04-14T03:27:06Z"
---
