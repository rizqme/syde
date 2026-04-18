---
id: REQ-0369
kind: requirement
name: Concept attributes field shall be removed
slug: concept-attributes-field-shall-be-removed-45o2
relationships:
    - target: entity-model
      type: refines
updated_at: "2026-04-18T09:37:56Z"
statement: The syde entity model shall not include an attributes field on concept entities.
req_type: constraint
priority: must
verification: No attributes field
source: plan
requirement_status: active
rationale: Properties in code
audited_overlaps:
    - slug: concept-actions-field-shall-be-removed
      distinction: attributes are typed properties while actions are domain verbs — different concept ERD elements removed
    - slug: concept-actions-field-shall-be-removed-6iik
      distinction: Removes the attributes field specifically; the actions field removal targets a different concept field name with independent schema and data-loss implications.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:37:56Z"
---
