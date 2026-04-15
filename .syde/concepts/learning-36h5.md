---
attributes:
    - description: LRN-NNNN counter ID
      name: id
    - description: short headline
      name: name
    - description: the actual learning content
      name: description
    - description: gotcha|constraint|pattern|anti-pattern|insight
      name: category
    - description: low|medium|high
      name: confidence
    - description: entity slugs this learning applies to
      name: entity_refs
description: Captured design knowledge — gotcha, constraint, convention, or note.
id: CPT-0008
invariants: category is one of the enumerated set. Confidence is high|medium|low. entity_refs must resolve.
kind: concept
lifecycle: remembered via 'syde remember' → optionally promoted to a Decision via 'syde learn promote'
meaning: A captured piece of design knowledge — gotcha, constraint, convention, context, dependency, performance note, or workaround
name: Learning
relationships:
    - target: syde
      type: belongs_to
    - target: entity-model
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: learning-36h5
structure_notes: Learning embeds BaseEntity plus category, entity_refs, confidence, source. Notes that don't fit into structured entity fields live here.
updated_at: "2026-04-14T10:48:03Z"
---
