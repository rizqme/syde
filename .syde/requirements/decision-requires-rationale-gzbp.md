---
id: REQ-0018
kind: requirement
name: Decision Requires Rationale
slug: decision-requires-rationale-gzbp
relationships:
    - target: decision-m2um
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:17Z"
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
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:17Z"
---
