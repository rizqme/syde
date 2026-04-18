---
id: REQ-0361
kind: requirement
name: Catch-all flow shall be replaced with user-goal flows
slug: catch-all-flow-shall-be-replaced-with-user-goal-flows-06lm
description: Design Model Operations Coverage deleted; user-goal flows created
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:05Z"
statement: The syde design model shall not contain a catch-all flow covering all contracts and shall instead use per-user-goal flows with structured steps.
req_type: constraint
priority: must
verification: syde query design-model-operations-coverage returns not found; all contracts covered by user-goal flow steps
source: plan
requirement_status: active
rationale: A single flow covering 71 contracts documents nothing; per-goal flows document everything
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:05Z"
---
