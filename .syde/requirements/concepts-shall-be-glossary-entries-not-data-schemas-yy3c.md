---
id: REQ-0368
kind: requirement
name: Concepts shall be glossary entries not data schemas
slug: concepts-shall-be-glossary-entries-not-data-schemas-yy3c
relationships:
    - target: entity-model
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T08:25:32Z"
statement: The syde entity model shall define concept entities as domain glossary entries with meaning, invariants, and lifecycle fields.
req_type: functional
priority: must
verification: ConceptEntity has meaning, invariants, lifecycle
source: plan
requirement_status: active
rationale: Concepts explain terms, not schemas
---
