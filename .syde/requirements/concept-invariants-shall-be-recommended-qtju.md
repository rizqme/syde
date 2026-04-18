---
id: REQ-0372
kind: requirement
name: Concept invariants shall be recommended
slug: concept-invariants-shall-be-recommended-qtju
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:17Z"
statement: The syde audit engine shall warn when a concept entity has an empty invariants field.
req_type: functional
priority: should
verification: sync check warns on empty invariants
source: plan
requirement_status: active
rationale: Invariants document rules
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:17Z"
---
