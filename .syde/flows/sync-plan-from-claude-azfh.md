---
id: FLW-0011
kind: flow
name: Sync Plan from Claude
slug: sync-plan-from-claude-azfh
description: Agent syncs a Claude Code plan file into syde plan format
tags:
    - planning
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: Agent has a Claude Code plan file to import
goal: syde plan entity matches the Claude plan content
steps:
    - id: s1
      action: Agent runs syde plan sync
      contract: sync-plan-from-claude
      description: Imports plan from Claude format
      on_success: s2
    - id: s2
      action: Agent lists tasks
      contract: list-tasks
      description: Verifies task mapping
      on_success: done
---
