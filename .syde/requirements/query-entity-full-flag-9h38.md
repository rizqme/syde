---
id: REQ-0290
kind: requirement
name: Query Entity Full Flag
slug: query-entity-full-flag-9h38
relationships:
    - target: query-entity-ci7d
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: Where --full is passed to syde query, the syde CLI shall include the entity body and all related data in the output.
req_type: interface
priority: must
verification: integration test invoking syde query --full
source: manual
source_ref: contract:query-entity-ci7d
requirement_status: active
rationale: The full flag unlocks the deep-dive view needed for targeted investigation.
---
