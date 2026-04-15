---
id: REQ-0077
kind: requirement
name: Relationship Target Resolution
slug: relationship-target-resolution-olj3
relationships:
    - target: relationship-hjgt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:54Z"
statement: If a relationship target slug or ID does not resolve to an existing entity at validation time, then the syde CLI shall reject the parent save with a validation error.
req_type: constraint
priority: must
verification: integration test creating a relationship to a nonexistent target
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: Unresolved targets break graph traversal and impact analysis.
---
