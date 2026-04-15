---
id: REQ-0092
kind: requirement
name: Relationship Requires Source And Target
slug: relationship-requires-source-and-target-n4bp
relationships:
    - target: relationship-hjgt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:02Z"
statement: The syde CLI shall require a non-empty source and target on every relationship instance.
req_type: constraint
priority: must
verification: unit test rejecting a relationship with an empty source or target
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: A relationship without both endpoints cannot form a graph edge.
---
