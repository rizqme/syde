---
id: CON-0016
kind: contract
name: Init Project
slug: init-project-xm6c
description: Scaffold .syde/ in the current project.
relationships:
    - target: syde-cli
      type: belongs_to
    - target: cli-commands
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: cli
interaction_pattern: request-response
input: syde init [--install-skill]
input_parameters:
    - path: --install-skill
      type: bool
      description: also run 'syde install-skill' after init
output: exit 0; creates .syde/ directory structure and syde.yaml
output_parameters:
    - path: created_dirs
      type: '[]string'
      description: list of created entity subdirectories
    - path: config_path
      type: string
      description: path to new syde.yaml
---
