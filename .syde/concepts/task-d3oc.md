---
attributes:
    - description: TSK-NNNN counter ID
      name: id
    - description: short task title
      name: name
    - description: parent plan slug
      name: plan
      refs:
        - plan
    - description: parent plan phase ID
      name: phase
      refs:
        - plan-phase
    - description: pending
      name: status
      refs:
        - in_progress|completed|blocked
    - description: what the task achieves
      name: objective
    - description: how to know it is done
      name: acceptance
    - description: existing entity slugs this task will modify
      name: affected_entities
      refs:
        - entity
    - description: concrete source file paths this task will touch
      name: affected_files
description: A tracked work item that references existing entities and files it will modify.
id: CPT-0003
invariants: affected_entities must all resolve. affected_files must all exist in the tree. task done on the last task in a phase auto-completes the phase.
kind: concept
lifecycle: pending → in_progress → completed | blocked | cancelled
meaning: A tracked work unit that references existing entities and files it will modify
name: Task
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: references
    - target: plan
      type: relates_to
slug: task-d3oc
structure_notes: Task embeds BaseEntity plus task_status, priority, objective, details, acceptance, plan_ref, plan_phase, affected_entities, affected_files, entity_refs (legacy). Task done auto-bumps updated_at on every affected entity.
updated_at: "2026-04-14T10:48:03Z"
---
