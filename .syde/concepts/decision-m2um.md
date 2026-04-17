---
id: CPT-0009
kind: concept
name: Decision
slug: decision-m2um
description: An architectural decision record (ADR) with rationale, tradeoffs, and consequences.
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: implemented_by
updated_at: "2026-04-17T08:25:52Z"
meaning: A durable architectural decision record (ADR)
invariants: Once approved, decisions should not be edited — create a superseding decision instead (manual convention, not enforced).
---
