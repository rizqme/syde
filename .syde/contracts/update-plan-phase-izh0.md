---
contract_kind: cli
description: Edit a plan phase's name, status, objective, changes, details, or notes.
id: CON-0028
input: syde plan phase <plan-slug> <phase-id> [flags]
input_parameters:
    - description: positional, required
      path: plan-slug
      type: string
    - description: positional, required. e.g. phase_1
      path: phase-id
      type: string
    - description: update name
      path: --name
      type: string
    - description: change parent phase ID
      path: --parent
      type: string
    - description: pending|in_progress|completed|skipped
      path: --status
      type: string
    - description: update description
      path: --description
      type: string
    - description: update objective
      path: --objective
      type: string
    - description: update changes
      path: --changes
      type: string
    - description: update details
      path: --details
      type: string
    - description: update notes
      path: --notes
      type: string
interaction_pattern: request-response
kind: contract
name: Update Plan Phase
output: exit 0; prints updated phase
output_parameters:
    - description: echoed phase id
      path: phase_id
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
slug: update-plan-phase-izh0
updated_at: "2026-04-14T03:27:04Z"
---
