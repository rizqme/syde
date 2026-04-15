---
contract_kind: cli
description: Return breadcrumb + summary + content for a tree node in one shot.
id: CON-0048
input: syde tree context <path> [--format plain|json] [--no-content] [--max-bytes N]
input_parameters:
    - description: positional, required
      path: path
      type: string
    - description: plain (default) or json
      path: --format
      type: string
    - description: omit file body
      path: --no-content
      type: bool
    - description: 'cap on inlined content (default: 64 KiB)'
      path: --max-bytes
      type: int
interaction_pattern: request-response
kind: contract
name: Tree Context Bundle
output: Breadcrumb + node summary + file content (or folder children listing)
output_parameters:
    - description: ancestor folder summaries from root to parent
      path: breadcrumb
      type: '[]string'
    - description: the node's own summary
      path: summary
      type: string
    - description: raw file bytes (files only, truncated to max-bytes)
      path: content
      type: string
    - description: direct children with summaries (folders only)
      path: children
      type: '[]string'
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
slug: tree-context-bundle-3co6
updated_at: "2026-04-14T03:27:05Z"
---
