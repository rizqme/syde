---
id: REQ-0074
kind: requirement
name: Plan Requires Objective
slug: plan-requires-objective-vacv
relationships:
    - target: plan-sk33
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:48Z"
statement: The syde CLI shall require a non-empty objective on every plan instance.
req_type: constraint
priority: must
verification: integration test running syde plan create without --objective
source: manual
source_ref: concept:plan-sk33
requirement_status: active
rationale: A plan without an objective cannot be evaluated for completion.
---
