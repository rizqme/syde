---
id: REQ-0352
kind: requirement
name: Audit shall error on contracts not in any flow step
slug: audit-shall-error-on-contracts-not-in-any-flow-step-2vck
description: ERROR when contract slug missing from all flow steps
relationships:
    - target: syde
      type: belongs_to
    - target: audit-engine
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: The syde audit engine shall report an error for any contract entity whose slug does not appear in the contract field of at least one flow step across all flows.
req_type: functional
priority: must
verification: syde sync check errors on contracts missing from all flow steps
source: plan
requirement_status: active
rationale: Every contract boundary must participate in a documented user journey
---
