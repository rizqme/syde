---
id: REQ-0396
kind: requirement
name: Requirement creation shall block unacknowledged overlaps
slug: requirement-creation-shall-block-unacknowledged-overlaps-dbzf
relationships:
    - target: cli-commands
      type: refines
updated_at: "2026-04-18T09:37:25Z"
statement: When syde add requirement detects one or more candidate overlaps above the audit threshold, the CLI shall refuse to create the entity unless every surfaced overlap is acknowledged with an --audited slug:reason entry or the author passes --force.
req_type: functional
priority: must
verification: syde add requirement with a high-overlap statement exits non-zero without --audited and succeeds with slug:reason
source: plan
requirement_status: active
rationale: Shift-left the overlap gate so authors resolve before the entity exists.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:25Z"
---
