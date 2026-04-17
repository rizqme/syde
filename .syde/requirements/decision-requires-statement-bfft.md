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
updated_at: "2026-04-17T10:46:09Z"
statement: The syde CLI shall require a statement on every decision instance.
req_type: constraint
priority: must
verification: integration test running syde add decision without --statement
source: manual
source_ref: concept:decision-m2um
requirement_status: active
rationale: A decision without a statement has no content to record.
audited_overlaps:
    - slug: decision-requires-category-5wfe
      distinction: Requires the statement field on decisions; the category requirement enforces a separate mandatory field for taxonomic classification, not the decision text.
    - slug: decision-requires-rationale
      distinction: statement captures what-was-decided whereas rationale captures why-the-decision — different attributes of the same entity kind
---
