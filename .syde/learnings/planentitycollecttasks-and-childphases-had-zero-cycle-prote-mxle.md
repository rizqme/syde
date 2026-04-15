---
category: gotcha
confidence: medium
description: 'PlanEntity.CollectTasks and ChildPhases had zero cycle protection. A corrupt plan with 5 phases all having empty IDs and empty ParentPhase caused ChildPhases('''') to match every phase (including itself), leading to infinite recursion and a 1GB goroutine stack dump. Fix: ChildPhases now skips empty-ID and self-parent phases; CollectTasks threads a visited map so any phase ID is walked at most once.'
discovered_at: "2026-04-14T09:06:18Z"
entity_refs:
    - entity-model
id: LRN-0008
kind: learning
name: PlanEntity.CollectTasks and ChildPhases had zero cycle prote
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: planentitycollecttasks-and-childphases-had-zero-cycle-prote-mxle
source: session-observation
updated_at: "2026-04-14T09:06:18Z"
---
