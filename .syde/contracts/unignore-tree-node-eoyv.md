---
contract_kind: cli
description: Remove the ignored flag from a tree node.
id: CON-0051
input: syde tree unignore <path>
input_parameters:
    - description: positional, required
      path: path
      type: string
interaction_pattern: request-response
kind: contract
name: Unignore Tree Node
output: exit 0; node un-ignored and marked stale for next summarize pass
output_parameters:
    - description: echoed un-ignored path
      path: confirmation
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: unignore-tree-node-eoyv
updated_at: "2026-04-14T03:27:05Z"
---
