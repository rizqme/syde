---
contract_kind: cli
description: Check the entire model for errors and drift warnings.
id: CON-0015
input: syde validate
input_parameters:
    - description: no arguments or flags
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Validate Model
output: exit 0 on success, 1 on errors. Prints errors and warnings grouped by severity
output_parameters:
    - description: fatal issues (missing required fields, unknown targets, cycles, orphan files, contract schema gaps)
      path: errors
      type: '[]string'
    - description: soft issues (drift between tree mtime and entity updated_at)
      path: warnings
      type: '[]string'
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - target: summary-tree
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: validate-model-tjzs
updated_at: "2026-04-14T03:27:04Z"
---
