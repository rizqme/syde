---
id: REQ-0280
kind: requirement
name: List Entities Invocation
slug: list-entities-invocation-bui6
relationships:
    - target: list-entities-0iec
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde list, the syde CLI shall print a tabular listing with one row per entity containing kind, name, and slug.
req_type: interface
priority: must
verification: integration test invoking syde list
source: manual
source_ref: contract:list-entities-0iec
requirement_status: active
rationale: Listing is the fastest way for operators to discover what exists in the model.
---
