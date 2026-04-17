---
id: FLW-0019
kind: flow
name: Check Constraints
slug: check-constraints-y0te
description: User checks which component and decisions apply to a file
tags:
    - querying
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User runs syde constraints check <file>
goal: User sees the owning component and applicable decisions for the file
steps:
    - id: s1
      action: User runs syde constraints check <file>
      contract: constraints-for-file
      description: CLI looks up file ownership
      on_success: s2
    - id: s2
      action: System shows applicable constraints
      contract: show-constraints
      description: Lists decisions for the component
      on_success: done
---
