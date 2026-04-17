---
id: FLW-0010
kind: flow
name: Update Plan
slug: update-plan-72e7
description: Agent modifies plan fields or phase details during execution
tags:
    - planning
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: Agent discovers the plan needs revision during implementation
goal: Plan reflects the revised approach
steps:
    - id: s1
      action: Agent updates plan header
      contract: update-plan
      description: Changes background/objective/scope/design
      on_success: s2
    - id: s2
      action: Agent updates phase details
      contract: update-plan-phase
      description: Changes phase objective/changes/details/notes
      on_success: s3
    - id: s3
      action: Agent links tasks to design entities
      contract: link-task-to-design
      description: Connects task to entity refs
      on_success: done
---
