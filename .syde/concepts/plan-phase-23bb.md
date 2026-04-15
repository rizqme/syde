---
attributes:
    - description: phase_N auto-generated within a plan
      name: id
    - description: short phase label, required
      name: name
    - description: optional parent phase ID for nesting
      name: parent_phase
      refs:
        - plan-phase
    - description: pending
      name: status
      refs:
        - in_progress|completed|skipped
    - description: what this phase delivers
      name: description
    - description: success condition
      name: objective
    - description: concrete list of what changes
      name: changes
    - description: implementation walkthrough
      name: details
    - description: risks / reminders / free-form context
      name: notes
    - description: task slugs linked to this phase
      name: tasks
      refs:
        - task
description: A milestone within a plan, holding objective/changes/details and a task list.
id: CPT-0007
invariants: parent_phase must resolve to another phase in the same plan. Phase ID is unique within plan.
kind: concept
lifecycle: pending → in_progress → completed | skipped. Auto-completes when all tasks are done.
meaning: One deliverable milestone within a plan, holding its own objective/changes/details and a list of tasks
name: Plan Phase
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: references
    - target: plan
      type: relates_to
slug: plan-phase-23bb
structure_notes: PlanPhase has id (phase_N), name, parent_phase, description, objective, changes, details, notes, status, and a Tasks slice with affected-entity/affected-file references.
updated_at: "2026-04-14T10:48:03Z"
---
