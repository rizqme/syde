---
id: CPT-0007
kind: concept
name: Plan Phase
slug: plan-phase-23bb
description: A milestone within a plan, holding objective/changes/details and a task list.
relationships:
    - target: syde
      type: belongs_to
    - target: plan
      type: relates_to
    - target: entity-model
      type: implemented_by
updated_at: "2026-04-17T08:25:52Z"
meaning: One deliverable milestone within a plan, holding its own objective/changes/details and a list of tasks
lifecycle: pending → in_progress → completed | skipped. Auto-completes when all tasks are done.
invariants: parent_phase must resolve to another phase in the same plan. Phase ID is unique within plan.
---
