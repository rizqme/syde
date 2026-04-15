---
attributes:
    - description: PLN-NNNN counter ID
      name: id
    - description: plan name
      name: name
    - description: why this plan exists
      name: background
    - description: what success looks like
      name: objective
    - description: in-scope / out-of-scope summary
      name: scope
    - description: draft
      name: plan_status
      refs:
        - approved|in-progress|completed
    - description: ordered hierarchical phase list
      name: phases
      refs:
        - plan-phase
    - description: set when plan is approved
      name: approved_at
    - description: set when all phases done
      name: completed_at
description: An implementation plan with background, objective, scope, phases, and tasks.
id: CPT-0002
invariants: Cannot execute until approved. Phase cannot complete until all its tasks are done. Parent phase cannot complete until all child phases are done.
kind: concept
lifecycle: draft → approved (explicit user chat approval) → in-progress → completed
meaning: An implementation plan with background, objective, scope, and a tree of phases containing tasks
name: Plan
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: references
    - target: entity
      type: relates_to
slug: plan-sk33
structure_notes: Plan embeds BaseEntity plus plan_status, background, objective, scope, phases[]. Phases nest via parent_phase. Each phase has its own objective/changes/details/notes and a task list.
updated_at: "2026-04-14T10:48:02Z"
---
