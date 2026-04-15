---
contract_kind: cli
description: Mark a task done; auto-bump updated_at on affected entities and parent phase.
id: CON-0038
input: syde task done <slug>
input_parameters:
    - description: positional, required
      path: slug
      type: string
interaction_pattern: request-response
kind: contract
name: Complete Task
output: exit 0; auto-bumps updated_at on affected entities and auto-completes parent phase if all tasks done
output_parameters:
    - description: completed
      path: new_status
      type: string
    - description: entities whose updated_at was bumped (clears drift warnings)
      path: touched_entities
      type: '[]string'
    - description: whether the parent phase was auto-completed
      path: phase_auto_completed
      type: bool
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
slug: complete-task-k8je
updated_at: "2026-04-14T03:27:04Z"
---
