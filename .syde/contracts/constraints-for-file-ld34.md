---
id: CON-0019
kind: contract
name: Constraints For File
slug: constraints-for-file-ld34
description: Map a source file to its component and report applicable constraints.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde constraints check <file> [--json]
input_parameters:
    - path: file
      type: string
      description: positional, required. Source file path
    - path: --json
      type: bool
      description: JSON output
output: Applicable decisions for the file's component
output_parameters:
    - path: component
      type: string
      description: matched component slug (from syde.yaml component_paths)
    - path: decisions
      type: '[]string'
      description: applicable architecture decisions
---
