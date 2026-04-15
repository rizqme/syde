---
id: REQ-0281
kind: requirement
name: List Entities Kind Filter
slug: list-entities-kind-filter-6c0r
relationships:
    - target: list-entities-0iec
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When syde list is invoked with a positional kind argument, the syde CLI shall filter the listing to entities of that kind.
req_type: interface
priority: must
verification: integration test invoking syde list component
source: manual
source_ref: contract:list-entities-0iec
requirement_status: active
rationale: Kind filtering narrows results in projects with many entities.
---
