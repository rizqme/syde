---
contract_kind: cli
description: List stale tree paths sorted deepest-first for leaves-first summarizing.
id: CON-0044
input: syde tree changes [--format plain|json] [--leaves-only]
input_parameters:
    - description: plain (default) or json
      path: --format
      type: string
    - description: hide stale folders whose descendants are still stale
      path: --leaves-only
      type: bool
interaction_pattern: request-response
kind: contract
name: List Tree Changes
output: Sorted list of stale paths, deepest first
output_parameters:
    - description: stale node paths with type (file|dir)
      path: paths
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
slug: list-tree-changes-mqhv
updated_at: "2026-04-14T03:27:05Z"
---
