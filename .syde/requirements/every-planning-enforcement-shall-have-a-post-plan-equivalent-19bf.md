---
id: REQ-0400
kind: requirement
name: Every planning enforcement shall have a post-plan equivalent
slug: every-planning-enforcement-shall-have-a-post-plan-equivalent-19bf
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:04Z"
statement: The syde audit engine shall maintain, for every finding category emitted by plan_authoring, an equivalent finding category emitted by the post-plan audit chain, so that an intent missed at planning is caught against actual state.
req_type: constraint
priority: must
verification: a Go table-driven test iterates the planning-post-plan parity registry and asserts both sides fire on the same crafted input
source: plan
requirement_status: active
rationale: Symmetric audits prevent one-sided gates from being evaded.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:04Z"
---
