---
id: REQ-0376
kind: requirement
name: Relates-to shall drop cardinality labels
slug: relates-to-shall-drop-cardinality-labels-agm8
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:48Z"
statement: The syde audit engine shall not validate cardinality labels on relates_to relationships between concept entities.
req_type: constraint
priority: must
verification: relates_to without cardinality succeeds
source: plan
requirement_status: active
rationale: Prose labels more natural
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:48Z"
---
