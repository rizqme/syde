---
id: REQ-0297
kind: requirement
name: Search Entities Invocation
slug: search-entities-invocation-hi8f
relationships:
    - target: search-entities-l5n1
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:52Z"
statement: When the user runs syde search <query>, the syde CLI shall return a ranked list of matching entity slugs with score and snippet.
req_type: interface
priority: must
verification: integration test invoking syde search with a known keyword
source: manual
source_ref: contract:search-entities-l5n1
requirement_status: active
rationale: Full-text search is the primary way operators discover entities by content rather than slug.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:52Z"
---
