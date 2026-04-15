---
category: pattern
confidence: medium
description: 'Dropping a field from a model struct is a soft-breaking change when gopkg.in/yaml.v3 is in lenient mode (default): existing YAML files with the orphan key still load (unknown keys ignored) and the key drops from disk on next save. No migration needed as long as the information isn''t load-bearing. Used when removing Type from ConceptAttribute — 11 existing concepts round-tripped cleanly with stale type: keys silently vanishing.'
discovered_at: "2026-04-14T10:51:57Z"
entity_refs:
    - entity-model
id: LRN-0020
kind: learning
name: Dropping a field from a model struct is a soft-breaking chan
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: dropping-a-field-from-a-model-struct-is-a-soft-breaking-chan-k029
source: session-observation
updated_at: "2026-04-14T10:51:57Z"
---
