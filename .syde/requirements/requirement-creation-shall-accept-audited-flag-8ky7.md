---
id: REQ-0381
kind: requirement
name: Requirement creation shall accept audited flag
slug: requirement-creation-shall-accept-audited-flag-8ky7
relationships:
    - target: cli-commands
      type: refines
updated_at: "2026-04-18T09:37:47Z"
statement: The syde add requirement command shall accept a repeatable --audited flag that acknowledges a specific overlap by slug.
req_type: functional
priority: must
verification: --audited <slug> consumed by syde add requirement; acknowledged overlaps are not re-warned.
source: plan
requirement_status: active
rationale: Authors need a way to proceed when overlap is intentional.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:47Z"
---
