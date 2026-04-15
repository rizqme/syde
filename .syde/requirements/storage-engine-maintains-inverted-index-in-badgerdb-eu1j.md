---
id: REQ-0177
kind: requirement
name: Storage Engine Maintains Inverted Index In BadgerDB
slug: storage-engine-maintains-inverted-index-in-badgerdb-eu1j
relationships:
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:13Z"
statement: The storage engine shall maintain a BadgerDB inverted index over entities, tags, words, and relationships for fast lookups.
req_type: functional
priority: must
verification: inspection of internal/storage/index.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: A rebuildable inverted index enables sub-linear search and traversal.
---
