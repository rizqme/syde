---
id: CPT-0002
kind: concept
name: Plan
slug: plan-sk33
description: An implementation plan with background, objective, scope, phases, and tasks.
relationships:
    - target: entity
      type: relates_to
    - target: entity-model
      type: implemented_by
    - target: syde
      type: belongs_to
    - target: plan-lifecycle
      type: used_in
updated_at: "2026-04-17T08:25:52Z"
meaning: An implementation plan with background, objective, scope, and a tree of phases containing tasks
lifecycle: draft → approved (explicit user chat approval) → in-progress → completed
invariants: Cannot execute until approved. Phase cannot complete until all its tasks are done. Parent phase cannot complete until all child phases are done.
---
