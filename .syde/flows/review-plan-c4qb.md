---
id: FLW-0007
kind: flow
name: Review Plan
slug: review-plan-c4qb
description: User reviews a plan in the dashboard or CLI before approving
tags:
    - planning
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-plan-lifecycle-flows
      type: references
updated_at: "2026-04-17T10:43:19Z"
trigger: Agent presents a plan for review
goal: User understands the plan's scope, design, changes, and tasks
steps:
    - id: s1
      action: Agent runs syde plan open
      contract: open-plan-in-dashboard
      description: Opens plan in dashboard tab
      on_success: s2
    - id: s2
      action: Dashboard renders plan detail
      contract: plan-view-screen
      description: Shows design + changes + tasks
      on_success: s3
    - id: s3
      action: Agent runs syde plan show
      contract: show-plan
      description: Prints plan summary
      on_success: s4
    - id: s4
      action: Dashboard renders plans inbox
      contract: plans-inbox-screen
      description: Plans list view
      on_success: s5
    - id: s5
      action: Agent runs syde plan list
      contract: list-plans
      description: Enumerates all plans in project
      on_success: s6
    - id: s6
      action: Agent runs syde plan check
      contract: check-plan
      description: Pre-approval plan authoring gate
      on_success: done
---
