---
id: REQ-0017
kind: requirement
name: Decision Requires Statement
slug: decision-requires-statement-bfft
relationships:
    - target: decision-m2um
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:51:45Z"
statement: The syde CLI shall require a statement on every decision instance.
req_type: constraint
priority: must
verification: integration test running syde add decision without --statement
source: manual
source_ref: concept:decision-m2um
requirement_status: active
rationale: A decision without a statement has no content to record.
---
