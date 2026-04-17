---
id: CON-0044
kind: contract
name: List Tree Changes
slug: list-tree-changes-mqhv
description: List stale tree paths sorted deepest-first for leaves-first summarizing.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: summary-tree
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde tree changes [--format plain|json] [--leaves-only]
input_parameters:
    - path: --format
      type: string
      description: plain (default) or json
    - path: --leaves-only
      type: bool
      description: hide stale folders whose descendants are still stale
output: Sorted list of stale paths, deepest first
output_parameters:
    - path: paths
      type: '[]string'
      description: stale node paths with type (file|dir)
---
