---
id: COM-0004
kind: component
name: Storage Engine
slug: storage-engine-ahgm
description: Markdown FileStore + BadgerDB index unified into one Store with counter-based IDs.
purpose: Persist entities as git-friendly markdown while keeping fast lookups through a rebuildable index
files:
    - internal/storage/store.go
    - internal/storage/filestore.go
    - internal/storage/index.go
    - internal/storage/indexer.go
    - internal/storage/serializer.go
    - internal/storage/counters.go
relationships:
    - target: existing-syde-model-baseline-hcvj
      type: references
      label: requirement
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:58:06Z"
responsibility: Unified Store over FileStore (markdown) and Index (BadgerDB)
capabilities:
    - Serialize/deserialize entities as YAML frontmatter + markdown body
    - Inverted index in BadgerDB (entities, tags, words, relationships)
    - Counter-based ID allocation (PFX-NNNN) persisted in BadgerDB
    - Full reindex from markdown files (idempotent)
    - Three-form slug addressing (full, bare, parent/child path)
boundaries: Does NOT own entity schema (delegates to model). Does NOT render results (delegates to query).
---
