---
id: REQ-0353
kind: requirement
name: Audit shall error on unresolvable step contract refs
slug: audit-shall-error-on-unresolvable-step-contract-refs-kw1j
description: ERROR when step contract slug does not resolve
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:59Z"
statement: If a flow step references a contract slug that does not resolve to an existing contract entity, then the syde audit engine shall report an error.
req_type: functional
priority: must
verification: Nonexistent contract slug causes syde sync check to error
source: plan
requirement_status: active
rationale: Dangling refs rot silently; catch them at audit time
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:59Z"
---
