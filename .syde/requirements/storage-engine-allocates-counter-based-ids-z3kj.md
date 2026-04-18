---
id: REQ-0178
kind: requirement
name: Storage Engine Allocates Counter Based IDs
slug: storage-engine-allocates-counter-based-ids-z3kj
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:36:58Z"
statement: When creating a new entity, the storage engine shall allocate a counter-based ID of the form PFX-NNNN by incrementing the per-kind counter in BadgerDB.
req_type: functional
priority: must
verification: unit test of NextID in internal/storage/counters.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Sequential per-kind IDs keep entity identifiers stable and rebuildable.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:36:58Z"
---
