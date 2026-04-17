---
id: CON-0047
kind: contract
name: Get Tree Summary
slug: get-tree-summary-2vyd
description: Print just one tree node's stored summary.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde tree get <path>
input_parameters:
    - path: path
      type: string
      description: positional, required
output: Just the node's summary text on stdout
output_parameters:
    - path: summary
      type: string
      description: stored summary
---
