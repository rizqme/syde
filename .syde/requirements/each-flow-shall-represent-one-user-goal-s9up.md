---
id: REQ-0348
kind: requirement
name: Each flow shall represent one user goal
slug: each-flow-shall-represent-one-user-goal-s9up
description: Flows are user journeys, one per goal
tags:
    - contract-coverage-reviewed
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:46Z"
statement: When creating a flow entity, the syde skill shall require that each flow represents a single user goal with a clear trigger and outcome.
req_type: constraint
priority: must
verification: Skill docs teach the one-flow-per-goal pattern
source: plan
requirement_status: active
rationale: Granular flows are reviewable and auditable; catch-all flows are not
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:46Z"
---
