---
id: CON-0015
kind: contract
name: Validate Model
slug: validate-model-tjzs
description: Check the entire model for errors and drift warnings.
relationships:
- target: cli-commands
  type: references
- target: summary-tree
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde validate
input_parameters:
- path: _
  type: '-'
  description: no arguments or flags
output: exit 0 on success, 1 on errors. Prints errors and warnings grouped by severity
output_parameters:
- path: errors
  type: '[]string'
  description: fatal issues (missing required fields, unknown targets, cycles, orphan files, contract schema gaps)
- path: warnings
  type: '[]string'
  description: soft issues (drift between tree mtime and entity updated_at)
---
