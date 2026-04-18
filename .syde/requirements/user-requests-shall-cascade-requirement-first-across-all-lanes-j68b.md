---
id: REQ-0339
kind: requirement
name: User requests shall cascade requirement-first across all lanes
slug: user-requests-shall-cascade-requirement-first-across-all-lanes-j68b
description: Planning workflow requirement that user-driven changes start with the Requirements lane.
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:56Z"
statement: When the user requests a change to the syde system, the agent shall first capture the underlying requirement as a New or Extended Requirement entry in the plan's Requirements lane before declaring any Component, Contract, Concept, or Flow changes that implement it.
req_type: constraint
priority: must
verification: 'Manual workflow inspection: every user-driven plan shows its Requirements lane populated before implementation lanes, and changed flows appear in the Flows lane.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: Requirements are the why, the other lanes are the how. The cascade rule forces explicit acknowledgement of changed behavior.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:56Z"
---
