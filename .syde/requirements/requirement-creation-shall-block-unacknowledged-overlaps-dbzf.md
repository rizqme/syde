---
id: REQ-0396
kind: requirement
name: Requirement creation shall block unacknowledged overlaps
slug: requirement-creation-shall-block-unacknowledged-overlaps-dbzf
relationships:
    - target: cli-commands
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:32Z"
statement: When syde add requirement detects one or more candidate overlaps above the audit threshold, the CLI shall refuse to create the entity unless every surfaced overlap is acknowledged with an --audited slug:reason entry or the author passes --force.
req_type: functional
priority: must
verification: syde add requirement with a high-overlap statement exits non-zero without --audited and succeeds with slug:reason
source: plan
requirement_status: active
rationale: Shift-left the overlap gate so authors resolve before the entity exists.
---
