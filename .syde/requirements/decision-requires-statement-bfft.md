---
id: REQ-0017
kind: requirement
name: Decision Requires Statement
slug: decision-requires-statement-bfft
relationships:
    - target: decision-m2um
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:28Z"
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
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:28Z"
---
