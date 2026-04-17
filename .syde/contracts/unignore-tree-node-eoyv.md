---
id: CON-0051
kind: contract
name: Unignore Tree Node
slug: unignore-tree-node-eoyv
description: Remove the ignored flag from a tree node.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde tree unignore <path>
input_parameters:
    - path: path
      type: string
      description: positional, required
output: exit 0; node un-ignored and marked stale for next summarize pass
output_parameters:
    - path: confirmation
      type: string
      description: echoed un-ignored path
---
