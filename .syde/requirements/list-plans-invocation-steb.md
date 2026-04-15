---
id: REQ-0282
kind: requirement
name: List Plans Invocation
slug: list-plans-invocation-steb
relationships:
    - target: list-plans-mo9k
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde plan list, the syde CLI shall print a tabular list of plans with slug, status, and completed-over-total task counts.
req_type: interface
priority: must
verification: integration test invoking syde plan list
source: manual
source_ref: contract:list-plans-mo9k
requirement_status: active
rationale: Plan listings give operators a quick progress dashboard across ongoing work.
---
