---
id: CON-0046
kind: contract
name: Show Tree
slug: show-tree-t0as
description: Render an ASCII tree with inline summaries and stale markers.
relationships:
- target: summary-tree
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde tree show [path] [--full] [--max-depth N] [--stale]
input_parameters:
- path: path
  type: string
  description: positional, optional. Root of the subtree to render
- path: --full
  type: bool
  description: unlimited depth
- path: --max-depth
  type: int
  description: 'depth cap (default: 2)'
- path: --stale
  type: bool
  description: prefix stale entries with '!'
output: ASCII tree with inline summaries
output_parameters:
- path: tree
  type: string
  description: rendered tree
---
