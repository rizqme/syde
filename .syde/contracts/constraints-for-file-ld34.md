---
contract_kind: cli
description: Map a source file to its component and report applicable constraints.
id: CON-0019
input: syde constraints check <file> [--json]
input_parameters:
    - description: positional, required. Source file path
      path: file
      type: string
    - description: JSON output
      path: --json
      type: bool
interaction_pattern: request-response
kind: contract
name: Constraints For File
output: Applicable decisions for the file's component
output_parameters:
    - description: matched component slug (from syde.yaml component_paths)
      path: component
      type: string
    - description: applicable architecture decisions
      path: decisions
      type: '[]string'
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: constraints-for-file-ld34
updated_at: "2026-04-14T03:27:04Z"
---
