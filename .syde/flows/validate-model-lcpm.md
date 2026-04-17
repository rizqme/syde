---
id: FLW-0020
kind: flow
name: Validate Model
slug: validate-model-lcpm
description: User runs the full health check on the design model
tags:
    - querying
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User runs syde sync check or syde validate
goal: User sees all errors, warnings, and hints across the model
steps:
    - id: s1
      action: User runs syde sync check
      contract: validate-model
      description: Runs audit.Run with all checks
      on_success: s2
    - id: s2
      action: System runs sync codebase check
      contract: sync-codebase
      description: Verifies model-code alignment
      on_success: done
---
