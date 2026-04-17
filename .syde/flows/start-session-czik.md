---
id: FLW-0032
kind: flow
name: Start Session
slug: start-session-czik
description: Agent session starts and loads architecture context
tags:
    - init
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: Claude Code session starts in a project with .syde/
goal: Agent has full architecture snapshot in context
steps:
    - id: s1
      action: SessionStart hook fires
      contract: session-context
      description: Runs syde context --json
      on_success: s2
    - id: s2
      action: Agent checks project status
      contract: project-status
      description: Verifies model health
      on_success: done
---
