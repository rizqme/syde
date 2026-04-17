---
id: FLW-0006
kind: flow
name: Create Plan
slug: create-plan-alr2
description: User creates a new plan with background, objective, scope, design, phases, and tasks
tags:
    - planning
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: User asks agent to implement a non-trivial change
goal: A draft plan with phases and tasks exists and is ready for review
steps:
    - id: s1
      action: User describes the change they want
      on_success: s2
    - id: s2
      action: Agent clarifies requirements
      on_success: s3
    - id: s3
      action: Agent runs syde plan create
      contract: create-plan
      description: Creates plan entity
      on_success: s4
    - id: s4
      action: Agent adds structured changes
      contract: add-plan-phase
      description: Adds phases
      on_success: s5
    - id: s5
      action: Agent creates tasks
      contract: create-task
      description: Creates tasks for each phase
      on_success: s6
    - id: s6
      action: Agent creates subtasks if needed
      contract: create-subtask
      description: Splits complex tasks
      on_success: s7
    - id: s7
      action: Agent estimates plan
      contract: estimate-plan
      description: Checks plan size
      on_success: s8
    - id: s8
      action: Agent runs syde plan check
      description: Validates plan authoring quality
      on_success: done
---
