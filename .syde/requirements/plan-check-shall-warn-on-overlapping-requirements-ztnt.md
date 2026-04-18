---
id: REQ-0344
kind: requirement
name: Plan check shall warn on overlapping requirements
slug: plan-check-shall-warn-on-overlapping-requirements-ztnt
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:45Z"
statement: When a new requirement in a plan shares significant terms with an existing active requirement, the syde audit engine shall warn that the requirements may overlap.
req_type: functional
priority: must
verification: A plan adding a requirement similar to an existing one triggers the overlap warning
source: plan
requirement_status: active
rationale: Overlapping requirements should be linked via refines/derives_from or the old one superseded
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:45Z"
---
