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
updated_at: "2026-04-17T10:46:09Z"
statement: The syde CLI shall require a rationale on every decision instance.
req_type: constraint
priority: must
verification: integration test running syde add decision without --rationale
source: manual
source_ref: concept:decision-m2um
requirement_status: active
rationale: ADR value comes from capturing why a decision was made.
audited_overlaps:
    - slug: decision-requires-category-5wfe
      distinction: Requires the rationale free-text field on decisions; the category requirement enforces a different mandatory field used to classify the decision into a taxonomy.
    - slug: decision-requires-statement
      distinction: rationale captures why-the-decision whereas statement captures what-was-decided — different attributes of the same entity kind
---
