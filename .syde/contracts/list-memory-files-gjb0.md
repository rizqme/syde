---
contract_kind: cli
description: List the Claude memory files written under .claude/.
id: CON-0066
input: syde memory list
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: List Memory Files
output: list of memory files under .claude/
output_parameters:
    - description: file paths
      path: rows
      type: '[]string'
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
slug: list-memory-files-gjb0
updated_at: "2026-04-14T03:27:06Z"
---
