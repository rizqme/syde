---
id: REQ-0291
kind: requirement
name: Query Entity Impacts Flag
slug: query-entity-impacts-flag-f98d
relationships:
    - target: query-entity-ci7d
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: Where --impacts is passed to syde query, the syde CLI shall return a transitive impact analysis starting at the named slug.
req_type: interface
priority: must
verification: integration test invoking syde query --impacts on a referenced entity
source: manual
source_ref: contract:query-entity-ci7d
requirement_status: active
rationale: Impact analysis supports change-risk reviews before editing architecturally central entities.
---
