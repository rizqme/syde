---
id: REQ-0370
kind: requirement
name: Concept actions field shall be removed
slug: concept-actions-field-shall-be-removed-6iik
relationships:
    - target: entity-model
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:50:15Z"
statement: The syde entity model shall not include an actions field on concept entities.
req_type: constraint
priority: must
verification: No actions field
source: plan
requirement_status: active
rationale: Verbs in lifecycle
audited_overlaps:
    - slug: concept-attributes-field-shall-be-removed
      distinction: actions are domain verbs while attributes are typed properties — different concept ERD elements removed
    - slug: concept-attributes-field-shall-be-removed-45o2
      distinction: Removes the actions field specifically; the attributes field removal is a separate model field with its own schema impact and migration.
---
