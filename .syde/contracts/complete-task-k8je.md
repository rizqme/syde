---
id: CON-0038
kind: contract
name: Complete Task
slug: complete-task-k8je
description: Mark a task done; auto-bump updated_at on affected entities and parent phase.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde task done <slug>
input_parameters:
    - path: slug
      type: string
      description: positional, required
output: exit 0; auto-bumps updated_at on affected entities and auto-completes parent phase if all tasks done
output_parameters:
    - path: new_status
      type: string
      description: completed
    - path: touched_entities
      type: '[]string'
      description: entities whose updated_at was bumped (clears drift warnings)
    - path: phase_auto_completed
      type: bool
      description: whether the parent phase was auto-completed
---
