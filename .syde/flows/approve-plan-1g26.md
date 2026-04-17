---
id: FLW-0008
kind: flow
name: Approve Plan
slug: approve-plan-1g26
description: User explicitly approves a plan, unlocking implementation
tags:
    - planning
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T09:12:26Z"
trigger: User says 'approve' in chat after reviewing the plan
goal: Plan transitions to approved status, writes are unblocked
steps:
    - id: s1
      action: User confirms approval in chat
      on_success: s2
    - id: s2
      action: Agent runs syde plan approve
      contract: approve-plan
      description: Transitions plan to approved, creates plan-sourced requirement
      on_success: done
---
