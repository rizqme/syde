---
contract_kind: cli
description: Print the full architecture snapshot used by the SessionStart hook.
id: CON-0017
input: syde context [--json]
input_parameters:
    - description: emit as JSON instead of markdown
      path: --json
      type: bool
interaction_pattern: request-response
kind: contract
name: Session Context
output: Full architecture snapshot printed on stdout (used by SessionStart hook)
output_parameters:
    - description: project metadata, entity counts, key decisions, top-level components, active plans
      path: snapshot
      type: string
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: session-context-74s7
updated_at: "2026-04-14T03:27:04Z"
---
