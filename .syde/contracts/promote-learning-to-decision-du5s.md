---
contract_kind: cli
description: Promote a learning into a formal decision record.
id: CON-0064
input: syde learn promote <slug> --to decision
input_parameters:
    - description: positional, required. Source learning slug
      path: slug
      type: string
    - description: target kind (currently only decision)
      path: --to
      type: string
interaction_pattern: request-response
kind: contract
name: Promote Learning To Decision
output: exit 0; prints new decision slug
output_parameters:
    - description: created decision slug
      path: new_slug
      type: string
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
slug: promote-learning-to-decision-du5s
updated_at: "2026-04-14T03:27:06Z"
---
