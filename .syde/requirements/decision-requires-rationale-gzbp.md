---
id: REQ-0018
kind: requirement
name: Decision Requires Rationale
slug: decision-requires-rationale-gzbp
relationships:
    - target: decision-m2um
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:51:46Z"
statement: The syde CLI shall require a rationale on every decision instance.
req_type: constraint
priority: must
verification: integration test running syde add decision without --rationale
source: manual
source_ref: concept:decision-m2um
requirement_status: active
rationale: ADR value comes from capturing why a decision was made.
---
