---
attributes:
    - description: DEC-NNNN counter ID
      name: id
    - description: short decision name
      name: name
    - description: the decision itself
      name: statement
    - description: why
      name: rationale
    - description: data|api|security|etc
      name: category
    - description: what else was weighed
      name: alternatives_considered
    - description: what we give up
      name: tradeoffs
    - description: downstream effects
      name: consequences
    - description: optional ID of decision this replaces
      name: supersedes
description: An architectural decision record (ADR) with rationale, tradeoffs, and consequences.
id: CPT-0009
invariants: Once approved, decisions should not be edited — create a superseding decision instead (manual convention, not enforced).
kind: concept
meaning: A durable architectural decision record (ADR)
name: Decision
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: references
slug: decision-m2um
structure_notes: Decision embeds BaseEntity plus category, statement, rationale, alternatives_considered, tradeoffs, consequences. Drives the constraints check.
updated_at: "2026-04-14T10:48:03Z"
---
