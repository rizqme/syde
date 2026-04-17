---
id: REQ-0019
kind: requirement
name: Decision Requires Category
slug: decision-requires-category-5wfe
relationships:
    - target: decision-m2um
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:09Z"
statement: The syde CLI shall require a category on every decision instance.
req_type: constraint
priority: must
verification: integration test running syde add decision without --category
source: manual
source_ref: concept:decision-m2um
requirement_status: active
rationale: Category enables constraint checks to filter decisions by domain.
audited_overlaps:
    - slug: decision-requires-rationale-gzbp
      distinction: Requires the category field on decisions; the rationale requirement enforces a different mandatory field capturing the why-of-the-decision free-text justification.
    - slug: decision-requires-statement-bfft
      distinction: Requires the category classification field on decisions; the statement requirement enforces a different mandatory field holding the decision statement text itself.
---
