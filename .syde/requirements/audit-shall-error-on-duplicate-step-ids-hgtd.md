---
id: REQ-0355
kind: requirement
name: Audit shall error on duplicate step IDs
slug: audit-shall-error-on-duplicate-step-ids-hgtd
description: ERROR when two steps in a flow share the same ID
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:47Z"
statement: If two or more steps within the same flow entity share the same id value, then the syde audit engine shall report an error.
req_type: functional
priority: must
verification: Two steps with id s1 in one flow causes syde sync check to error
source: plan
requirement_status: active
rationale: Step IDs are the addressing mechanism for the intra-flow graph
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:47Z"
---
