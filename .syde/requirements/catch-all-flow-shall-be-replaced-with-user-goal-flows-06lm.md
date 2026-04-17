---
id: REQ-0361
kind: requirement
name: Catch-all flow shall be replaced with user-goal flows
slug: catch-all-flow-shall-be-replaced-with-user-goal-flows-06lm
description: Design Model Operations Coverage deleted; user-goal flows created
relationships:
    - target: syde
      type: belongs_to
updated_at: "2026-04-16T10:40:59Z"
statement: The syde design model shall not contain a catch-all flow covering all contracts and shall instead use per-user-goal flows with structured steps.
req_type: constraint
priority: must
verification: syde query design-model-operations-coverage returns not found; all contracts covered by user-goal flow steps
source: plan
requirement_status: active
rationale: A single flow covering 71 contracts documents nothing; per-goal flows document everything
---
