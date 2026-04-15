---
contract_kind: cli
description: Delete stale Claude memory files.
id: CON-0067
input: syde memory clean
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Clean Memory Files
output: Deletes stale memory files
output_parameters:
    - description: number of removed files
      path: deleted
      type: int
relationships:
    - target: syde-cli
      type: belongs_to
    - target: memory-sync
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: clean-memory-files-i8h3
updated_at: "2026-04-14T03:27:06Z"
---
