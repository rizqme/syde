---
id: REQ-0179
kind: requirement
name: Storage Engine Rebuilds Index From Files
slug: storage-engine-rebuilds-index-from-files-xa7m
relationships:
    - target: storage-engine-ahgm
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:56:18Z"
statement: When Reindex is invoked, the storage engine shall scan every markdown file via the FileStore and rebuild the BadgerDB index from scratch.
req_type: functional
priority: must
verification: integration test of Reindex in internal/storage/indexer.go
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: A rebuildable index means files remain the source of truth.
---
