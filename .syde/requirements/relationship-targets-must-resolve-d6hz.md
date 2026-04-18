---
id: REQ-0023
kind: requirement
name: Relationship Targets Must Resolve
slug: relationship-targets-must-resolve-d6hz
relationships:
    - target: entity-8x6p
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:00Z"
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
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:00Z"
---
