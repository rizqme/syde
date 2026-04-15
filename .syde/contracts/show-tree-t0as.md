---
contract_kind: cli
description: Render an ASCII tree with inline summaries and stale markers.
id: CON-0046
input: syde tree show [path] [--full] [--max-depth N] [--stale]
input_parameters:
    - description: positional, optional. Root of the subtree to render
      path: path
      type: string
    - description: unlimited depth
      path: --full
      type: bool
    - description: 'depth cap (default: 2)'
      path: --max-depth
      type: int
    - description: prefix stale entries with '!'
      path: --stale
      type: bool
interaction_pattern: request-response
kind: contract
name: Show Tree
output: ASCII tree with inline summaries
output_parameters:
    - description: rendered tree
      path: tree
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: show-tree-t0as
updated_at: "2026-04-14T03:27:05Z"
---
