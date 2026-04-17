---
id: REQ-0348
kind: requirement
name: Each flow shall represent one user goal
slug: each-flow-shall-represent-one-user-goal-s9up
description: Flows are user journeys, one per goal
tags:
    - contract-coverage-reviewed
relationships:
    - target: syde
      type: belongs_to
updated_at: "2026-04-17T10:55:11Z"
statement: When creating a flow entity, the syde skill shall require that each flow represents a single user goal with a clear trigger and outcome.
req_type: constraint
priority: must
verification: Skill docs teach the one-flow-per-goal pattern
source: plan
requirement_status: active
rationale: Granular flows are reviewable and auditable; catch-all flows are not
---
