---
id: REQ-0179
kind: requirement
name: Storage Engine Rebuilds Index From Files
slug: storage-engine-rebuilds-index-from-files-xa7m
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:41Z"
statement: When Reindex is invoked, the storage engine shall scan every markdown file via the FileStore and rebuild the BadgerDB index from scratch.
req_type: functional
priority: must
verification: integration test of Reindex in internal/storage/indexer.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: A rebuildable index means files remain the source of truth.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:41Z"
---
