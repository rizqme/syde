---
id: CPT-0003
kind: concept
name: Task
slug: task-d3oc
description: A tracked work item that references existing entities and files it will modify.
relationships:
    - target: syde
      type: belongs_to
    - target: plan
      type: relates_to
    - target: entity-model
      type: implemented_by
updated_at: "2026-04-17T08:25:52Z"
meaning: A tracked work unit that references existing entities and files it will modify
lifecycle: pending → in_progress → completed | blocked | cancelled
invariants: affected_entities must all resolve. affected_files must all exist in the tree. task done on the last task in a phase auto-completes the phase.
---
