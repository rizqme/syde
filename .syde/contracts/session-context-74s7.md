---
id: CON-0017
kind: contract
name: Session Context
slug: session-context-74s7
description: Print the full architecture snapshot used by the SessionStart hook.
relationships:
- target: cli-commands
  type: references
- type: belongs_to
  target: syde-5tdt
updated_at: '2026-04-16T10:51:16Z'
contract_kind: cli
interaction_pattern: request-response
input: syde context [--json]
input_parameters:
- path: --json
  type: bool
  description: emit as JSON instead of markdown
output: Full architecture snapshot printed on stdout (used by SessionStart hook)
output_parameters:
- path: snapshot
  type: string
  description: project metadata, entity counts, key decisions, top-level components, active plans
---
