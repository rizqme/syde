---
contract_kind: cli
description: Print just one tree node's stored summary.
id: CON-0047
input: syde tree get <path>
input_parameters:
    - description: positional, required
      path: path
      type: string
interaction_pattern: request-response
kind: contract
name: Get Tree Summary
output: Just the node's summary text on stdout
output_parameters:
    - description: stored summary
      path: summary
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: get-tree-summary-2vyd
updated_at: "2026-04-14T03:27:05Z"
---
