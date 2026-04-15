---
contract_kind: cli
description: Flag a tree node as ignored so the orphan-file validator skips it.
id: CON-0050
input: syde tree ignore <path>
input_parameters:
    - description: positional, required
      path: path
      type: string
interaction_pattern: request-response
kind: contract
name: Ignore Tree Node
output: exit 0; node flagged as ignored (exempt from orphan validator)
output_parameters:
    - description: echoed ignored path
      path: confirmation
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: ignore-tree-node-m460
updated_at: "2026-04-14T03:27:05Z"
---
