---
id: REQ-0328
kind: requirement
name: Slug index key shall map slugs to entity IDs
slug: slug-index-key-shall-map-slugs-to-entity-ids-ja4z
relationships:
    - target: slug-index-key-m5af
      type: refines
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T11:03:44Z"
statement: The syde storage layer shall store each entity's slug-to-ID mapping under BadgerDB key 's:<kind>:<slug>'.
req_type: interface
priority: must
verification: Integration test creating an entity and resolving via slug
source: manual
source_ref: contract:slug-index-key-m5af
requirement_status: active
rationale: Slug index keys make slug-based CLI lookups a single BadgerDB hit instead of a file scan.
---
