---
contract_kind: cli
description: Show all active decisions plus high-confidence learnings.
id: CON-0068
input: syde constraints [--json]
input_parameters:
    - description: JSON output
      path: --json
      type: bool
interaction_pattern: request-response
kind: contract
name: Show Constraints
output: All active decisions + critical learnings
output_parameters:
    - description: all architecture decisions
      path: decisions
      type: '[]object'
    - description: high-confidence learnings
      path: learnings
      type: '[]object'
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
slug: show-constraints-jnig
updated_at: "2026-04-14T03:27:06Z"
---
