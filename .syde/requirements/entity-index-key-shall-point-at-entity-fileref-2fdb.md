---
id: REQ-0325
kind: requirement
name: Entity index key shall point at entity FileRef
slug: entity-index-key-shall-point-at-entity-fileref-2fdb
relationships:
    - target: entity-index-key-i12k
      type: refines
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T11:03:44Z"
statement: The syde storage layer shall store each entity's file path and frontmatter line span under BadgerDB key 'e:<kind>:<id>'.
req_type: interface
priority: must
verification: 'Integration test creating an entity and inspecting e: key contents'
source: manual
source_ref: contract:entity-index-key-i12k
requirement_status: active
rationale: Entity index keys let lookups return a file pointer without reading every markdown file.
---
