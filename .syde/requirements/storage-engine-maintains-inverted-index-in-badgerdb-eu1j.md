---
id: REQ-0177
kind: requirement
name: Storage Engine Maintains Inverted Index In BadgerDB
slug: storage-engine-maintains-inverted-index-in-badgerdb-eu1j
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:17Z"
statement: The storage engine shall maintain a BadgerDB inverted index over entities, tags, words, and relationships for fast lookups.
req_type: functional
priority: must
verification: inspection of internal/storage/index.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: A rebuildable inverted index enables sub-linear search and traversal.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:17Z"
---
