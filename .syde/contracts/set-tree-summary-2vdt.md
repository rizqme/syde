---
id: CON-0045
kind: contract
name: Set Tree Summary
slug: set-tree-summary-2vdt
description: Store a node's summary text and mark its parent stale.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde tree summarize <path> --summary ...
input_parameters:
    - path: path
      type: string
      description: positional, required. Relative path to a tracked node
    - path: --summary
      type: string
      description: required. Summary text; pass '-' to read from stdin
output: exit 0; prints 'Summarized <path> (parent <p> marked stale)'
output_parameters:
    - path: confirmation
      type: string
      description: human confirmation
---
