---
id: FLW-0030
kind: flow
name: Initialize Project
slug: initialize-project-5knu
description: User bootstraps a new syde project
tags:
    - init
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: User runs syde init in a project directory
goal: A .syde/ directory exists with config, index, and optional skill files
steps:
    - id: s1
      action: User runs syde init
      contract: init-project
      description: Creates .syde/ skeleton, syde.yaml, index
      on_success: s2
    - id: s2
      action: User runs syde install-skill
      contract: install-skill
      description: Writes SKILL.md, hooks, references to .claude/
      on_success: done
---
