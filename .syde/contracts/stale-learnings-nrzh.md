---
contract_kind: cli
description: List learnings whose linked entities have changed.
id: CON-0063
input: syde learn stale
input_parameters:
    - description: no arguments
      path: _
      type: '-'
interaction_pattern: request-response
kind: contract
name: Stale Learnings
output: learnings referencing changed entities
output_parameters:
    - description: stale learnings
      path: rows
      type: '[]string'
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: stale-learnings-nrzh
updated_at: "2026-04-14T03:27:06Z"
---
