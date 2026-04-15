---
id: REQ-0082
kind: requirement
name: Relationship Type Must Be Allowed
slug: relationship-type-must-be-allowed-8aar
relationships:
    - target: relationship-hjgt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:01Z"
statement: The syde CLI shall accept a relationship type only if it is one of belongs_to, depends_on, exposes, consumes, uses, involves, references, relates_to, implements, applies_to, modifies, or visualizes.
req_type: constraint
priority: must
verification: integration test adding a relationship with an unknown type
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: A closed type set keeps graph semantics interoperable across tooling.
---
