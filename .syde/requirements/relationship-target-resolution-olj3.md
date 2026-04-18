---
id: REQ-0077
kind: requirement
name: Relationship Target Resolution
slug: relationship-target-resolution-olj3
relationships:
    - target: relationship-hjgt
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:42Z"
statement: If a relationship target slug or ID does not resolve to an existing entity at validation time, then the syde CLI shall reject the parent save with a validation error.
req_type: constraint
priority: must
verification: integration test creating a relationship to a nonexistent target
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: Unresolved targets break graph traversal and impact analysis.
audited_overlaps:
    - slug: relationship-targets-must-resolve-d6hz
      distinction: This requirement scopes to the validation-time resolution pass; the paired one scopes to the general relationship declaration save.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:42Z"
---
