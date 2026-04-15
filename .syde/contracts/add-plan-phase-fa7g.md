---
contract_kind: cli
description: Add a phase (optionally nested) to an existing plan.
id: CON-0027
input: syde plan add-phase <plan-slug> [flags]
input_parameters:
    - description: positional, required
      path: plan-slug
      type: string
    - description: short phase label
      path: --name
      type: string
    - description: parent phase ID for nesting
      path: --parent
      type: string
    - description: phase description
      path: --description
      type: string
    - description: what this phase achieves
      path: --objective
      type: string
    - description: what concretely changes
      path: --changes
      type: string
    - description: implementation walkthrough
      path: --details
      type: string
    - description: risks, reminders, new entities to be created
      path: --notes
      type: string
interaction_pattern: request-response
kind: contract
name: Add Plan Phase
output: exit 0; prints new phase ID
output_parameters:
    - description: auto-generated phase ID (phase_N)
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
slug: add-plan-phase-fa7g
updated_at: "2026-04-14T03:27:04Z"
---
