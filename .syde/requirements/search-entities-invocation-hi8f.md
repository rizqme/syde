---
id: REQ-0297
kind: requirement
name: Search Entities Invocation
slug: search-entities-invocation-hi8f
relationships:
    - target: search-entities-l5n1
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde search <query>, the syde CLI shall return a ranked list of matching entity slugs with score and snippet.
req_type: interface
priority: must
verification: integration test invoking syde search with a known keyword
source: manual
source_ref: contract:search-entities-l5n1
requirement_status: active
rationale: Full-text search is the primary way operators discover entities by content rather than slug.
---
