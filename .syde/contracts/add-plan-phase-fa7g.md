---
id: CON-0027
kind: contract
name: Add Plan Phase
slug: add-plan-phase-fa7g
description: Add a phase (optionally nested) to an existing plan.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde plan add-phase <plan-slug> [flags]
input_parameters:
    - path: plan-slug
      type: string
      description: positional, required
    - path: --name
      type: string
      description: short phase label
    - path: --parent
      type: string
      description: parent phase ID for nesting
    - path: --description
      type: string
      description: phase description
    - path: --objective
      type: string
      description: what this phase achieves
    - path: --changes
      type: string
      description: what concretely changes
    - path: --details
      type: string
      description: implementation walkthrough
    - path: --notes
      type: string
      description: risks, reminders, new entities to be created
output: exit 0; prints new phase ID
output_parameters:
    - path: phase_id
      type: string
      description: auto-generated phase ID (phase_N)
---
