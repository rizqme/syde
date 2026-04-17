---
id: REQ-0400
kind: requirement
name: Every planning enforcement shall have a post-plan equivalent
slug: every-planning-enforcement-shall-have-a-post-plan-equivalent-19bf
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:32Z"
statement: The syde audit engine shall maintain, for every finding category emitted by plan_authoring, an equivalent finding category emitted by the post-plan audit chain, so that an intent missed at planning is caught against actual state.
req_type: constraint
priority: must
verification: a Go table-driven test iterates the planning-post-plan parity registry and asserts both sides fire on the same crafted input
source: plan
requirement_status: active
rationale: Symmetric audits prevent one-sided gates from being evaded.
---
