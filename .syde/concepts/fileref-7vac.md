---
id: CPT-0006
kind: concept
name: FileRef
slug: fileref-7vac
description: BadgerDB-side pointer back to a markdown file plus per-field line ranges.
relationships:
    - target: syde
      type: belongs_to
    - target: storage-engine
      type: implemented_by
updated_at: "2026-04-17T08:25:52Z"
meaning: An index-side pointer from BadgerDB back to a markdown file + line range
invariants: File must exist on disk when the index is queried. Line ranges must be valid positions within the file. Always rebuildable from markdown via 'syde reindex'.
---
