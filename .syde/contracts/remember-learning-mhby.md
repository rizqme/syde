---
contract_kind: cli
description: Capture an informal learning with category, confidence, and entity links.
id: CON-0022
input: syde remember <text> [flags]
input_parameters:
    - description: positional, required. The learning content
      path: text
      type: string
    - description: gotcha|constraint|convention|context|dependency|performance|workaround
      path: --category
      type: string
    - description: repeatable linked entity slug
      path: --entity
      type: '[]string'
    - description: 'high|medium|low (default: high)'
      path: --confidence
      type: string
interaction_pattern: request-response
kind: contract
name: Remember Learning
output: exit 0; prints created learning slug
output_parameters:
    - description: new learning slug
      path: slug
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
slug: remember-learning-mhby
updated_at: "2026-04-14T03:27:04Z"
---
