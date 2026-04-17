---
id: CON-0050
kind: contract
name: Ignore Tree Node
slug: ignore-tree-node-m460
description: Flag a tree node as ignored so the orphan-file validator skips it.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde tree ignore <path>
input_parameters:
    - path: path
      type: string
      description: positional, required
output: exit 0; node flagged as ignored (exempt from orphan validator)
output_parameters:
    - path: confirmation
      type: string
      description: echoed ignored path
---
