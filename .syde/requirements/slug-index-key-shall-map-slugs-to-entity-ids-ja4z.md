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
updated_at: "2026-04-18T09:37:52Z"
statement: The syde storage layer shall store each entity's slug-to-ID mapping under BadgerDB key 's:<kind>:<slug>'.
req_type: interface
priority: must
verification: Integration test creating an entity and resolving via slug
source: manual
source_ref: contract:slug-index-key-m5af
requirement_status: active
rationale: Slug index keys make slug-based CLI lookups a single BadgerDB hit instead of a file scan.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:52Z"
---
