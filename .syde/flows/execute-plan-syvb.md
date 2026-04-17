---
id: FLW-0009
kind: flow
name: Execute Plan
slug: execute-plan-syvb
description: Agent implements an approved plan phase by phase, task by task
tags:
    - planning
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T10:43:19Z"
trigger: Plan is approved and agent begins implementation
goal: All tasks completed, all phases auto-completed
steps:
    - id: s1
      action: Agent starts a task
      contract: start-task
      description: Transitions task to in_progress
      on_success: s2
    - id: s2
      action: Agent writes code
      on_success: s3
      on_failure: s5
    - id: s3
      action: Agent completes task
      contract: complete-task
      description: Marks done with affected entities/files
      on_success: s4
    - id: s4
      action: System auto-completes phase if all tasks done
      description: Phase status check
      on_success: s1
      on_failure: done
    - id: s5
      action: Agent blocks task if stuck
      contract: block-task
      description: Records block reason
      on_success: s6
    - id: s6
      action: Agent updates task details
      contract: update-task
      description: Adjusts objective/details/acceptance
      on_success: s1
    - id: s7
      action: Agent runs plan execute for scaffolding
      contract: execute-plan
      description: Scaffolds entity files
      on_success: s1
    - id: s8
      action: Agent runs sync check gate
      contract: sync-check-gate
      description: Canonical session-end health check
      on_success: s9
    - id: s9
      action: Agent runs plan complete
      contract: complete-plan
      description: Marks plan completed when all changes realize
      on_success: done
---
