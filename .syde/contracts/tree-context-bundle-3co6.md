---
id: CON-0048
kind: contract
name: Tree Context Bundle
slug: tree-context-bundle-3co6
description: Return breadcrumb + summary + content for a tree node in one shot.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
updated_at: "2026-04-16T10:51:16Z"
contract_kind: cli
interaction_pattern: request-response
input: syde tree context <path> [--format plain|json] [--no-content] [--max-bytes N]
input_parameters:
    - path: path
      type: string
      description: positional, required
    - path: --format
      type: string
      description: plain (default) or json
    - path: --no-content
      type: bool
      description: omit file body
    - path: --max-bytes
      type: int
      description: 'cap on inlined content (default: 64 KiB)'
output: Breadcrumb + node summary + file content (or folder children listing)
output_parameters:
    - path: breadcrumb
      type: '[]string'
      description: ancestor folder summaries from root to parent
    - path: summary
      type: string
      description: the node's own summary
    - path: content
      type: string
      description: raw file bytes (files only, truncated to max-bytes)
    - path: children
      type: '[]string'
      description: direct children with summaries (folders only)
---
