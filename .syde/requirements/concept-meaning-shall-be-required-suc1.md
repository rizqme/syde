---
id: REQ-0371
kind: requirement
name: Concept meaning shall be required
slug: concept-meaning-shall-be-required-suc1
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:39Z"
statement: The syde audit engine shall report an error for any concept entity with an empty meaning field.
req_type: functional
priority: must
verification: sync check errors on empty meaning
source: plan
requirement_status: active
rationale: Meaning is core
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:39Z"
---
