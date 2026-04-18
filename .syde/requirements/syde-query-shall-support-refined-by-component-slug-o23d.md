---
id: REQ-0483
kind: requirement
name: syde query shall support --refined-by component slug
slug: syde-query-shall-support-refined-by-component-slug-o23d
relationships:
    - target: query-engine-9k84
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:40Z"
statement: The syde query command shall support a --refined-by component-slug flag that lists every active requirement carrying a refines edge to the resolved component.
req_type: functional
priority: should
verification: syde query --refined-by <component-slug> returns the slugs of every active requirement with refines edges resolving to that component
source: plan
source_ref: bidirectional-requirement-component-coupling-with-content-hash-recheck-gate-p77e
requirement_status: active
rationale: Mirrors --depended-by/--depends-on naming; resolves component slug via the same alias map the audit uses.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:40Z"
    query-engine-9k84:
        hash: 03a24974e906ccbc86ac65d8d2da018434bef5290e59b82647d94ff0290ac1d3
        at: "2026-04-18T09:36:40Z"
---
