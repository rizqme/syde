---
id: REQ-0381
kind: requirement
name: Requirement creation shall accept audited flag
slug: requirement-creation-shall-accept-audited-flag-8ky7
relationships:
    - target: cli-commands
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T09:07:56Z"
statement: The syde add requirement command shall accept a repeatable --audited flag that acknowledges a specific overlap by slug.
req_type: functional
priority: must
verification: --audited <slug> consumed by syde add requirement; acknowledged overlaps are not re-warned.
source: plan
requirement_status: active
rationale: Authors need a way to proceed when overlap is intentional.
---
