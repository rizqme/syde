---
category: pattern
confidence: medium
description: Concept entities are now first-class ERD nodes with structured Attributes/Actions slices and a cardinality enum on relates_to ({one-to-one, one-to-many, many-to-one, many-to-many}). Validator requires meaning + >=1 attribute (each with non-empty name+type); audit rejects invalid cardinality labels. Bumping this is breaking — existing concepts without attributes turn into sync-check ERRORs and must be back-filled.
discovered_at: "2026-04-14T10:12:28Z"
entity_refs:
    - entity-model
id: LRN-0014
kind: learning
name: Concept entities are now first-class ERD nodes with structur
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: concept-entities-are-now-first-class-erd-nodes-with-structur-cuwn
source: session-observation
updated_at: "2026-04-14T10:12:28Z"
---
