---
id: REQ-0178
kind: requirement
name: Storage Engine Allocates Counter Based IDs
slug: storage-engine-allocates-counter-based-ids-z3kj
relationships:
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:15Z"
statement: When creating a new entity, the storage engine shall allocate a counter-based ID of the form PFX-NNNN by incrementing the per-kind counter in BadgerDB.
req_type: functional
priority: must
verification: unit test of NextID in internal/storage/counters.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Sequential per-kind IDs keep entity identifiers stable and rebuildable.
---
