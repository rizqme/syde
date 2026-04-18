---
id: REQ-0403
kind: requirement
name: Plan diffs introducing contracts shall also introduce matching flow steps
slug: plan-diffs-introducing-contracts-shall-also-introduce-matching-flow-steps-s8at
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:50Z"
statement: When a syde plan change diff adds or extends a contract entry, the same diff shall pair that entry with a flow lane item whose trajectory names the contract slug.
req_type: functional
priority: must
verification: plan check on a diff that adds a contract without a paired flow lane item emits a warning; pairing clears it
source: plan
requirement_status: active
rationale: Enforces co-evolution of boundary definitions and their traversal documents at planning time.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:50Z"
---
