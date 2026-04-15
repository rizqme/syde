---
contract_kind: cli
description: Store a node's summary text and mark its parent stale.
id: CON-0045
input: syde tree summarize <path> --summary ...
input_parameters:
    - description: positional, required. Relative path to a tracked node
      path: path
      type: string
    - description: required. Summary text; pass '-' to read from stdin
      path: --summary
      type: string
interaction_pattern: request-response
kind: contract
name: Set Tree Summary
output: exit 0; prints 'Summarized <path> (parent <p> marked stale)'
output_parameters:
    - description: human confirmation
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
slug: set-tree-summary-2vdt
updated_at: "2026-04-14T03:27:05Z"
---
