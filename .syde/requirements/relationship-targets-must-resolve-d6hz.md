---
id: REQ-0023
kind: requirement
name: Relationship Targets Must Resolve
slug: relationship-targets-must-resolve-d6hz
relationships:
    - target: entity-8x6p
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:07Z"
statement: If an entity declares a relationship whose target does not resolve to an existing entity, then the syde CLI shall reject the save with a validation error.
req_type: constraint
priority: must
verification: integration test adding a relationship to a non-existent slug
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Dangling relationships corrupt the graph used for impact analysis.
audited_overlaps:
    - slug: relationship-target-resolution-olj3
      distinction: Save-time relationship rejection differs from the validation-time target resolution rule covering slug and ID lookup.
---
