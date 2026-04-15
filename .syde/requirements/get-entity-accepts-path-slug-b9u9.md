---
id: REQ-0273
kind: requirement
name: Get Entity Accepts Path Slug
slug: get-entity-accepts-path-slug-b9u9
relationships:
    - target: get-entity-0lzq
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When syde get is invoked, the syde CLI shall accept slug as a full slug, bare slug, or parent/child path string.
req_type: interface
priority: must
verification: integration test invoking syde get with each slug form
source: manual
source_ref: contract:get-entity-0lzq
requirement_status: active
rationale: Flexible slug resolution reduces friction when operators remember only part of an identifier.
---
