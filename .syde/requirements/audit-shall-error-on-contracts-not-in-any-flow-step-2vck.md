---
id: REQ-0352
kind: requirement
name: Audit shall error on contracts not in any flow step
slug: audit-shall-error-on-contracts-not-in-any-flow-step-2vck
description: ERROR when contract slug missing from all flow steps
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:36:53Z"
statement: The syde audit engine shall report an error for any contract entity whose slug does not appear in the contract field of at least one flow step across all flows.
req_type: functional
priority: must
verification: syde sync check errors on contracts missing from all flow steps
source: plan
requirement_status: active
rationale: Every contract boundary must participate in a documented user journey
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:53Z"
---
