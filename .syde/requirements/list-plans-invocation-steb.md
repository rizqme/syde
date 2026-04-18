---
id: REQ-0282
kind: requirement
name: List Plans Invocation
slug: list-plans-invocation-steb
relationships:
    - target: list-plans-mo9k
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:11Z"
statement: When the user runs syde plan list, the syde CLI shall print a tabular list of plans with slug, status, and completed-over-total task counts.
req_type: interface
priority: must
verification: integration test invoking syde plan list
source: manual
source_ref: contract:list-plans-mo9k
requirement_status: active
rationale: Plan listings give operators a quick progress dashboard across ongoing work.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:11Z"
---
