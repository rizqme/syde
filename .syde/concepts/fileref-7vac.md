---
attributes:
    - description: relative path to the markdown file under .syde/
      name: file
    - description: entity ID the ref points at
      name: id
    - description: entity name copy for fast render
      name: name
    - description: entity kind
      name: kind
    - description: field name to line span in the markdown
      name: lines
data_sensitivity: Non-sensitive cache data. Safe to delete — rebuildable via reindex.
description: BadgerDB-side pointer back to a markdown file plus per-field line ranges.
id: CPT-0006
invariants: File must exist on disk when the index is queried. Line ranges must be valid positions within the file. Always rebuildable from markdown via 'syde reindex'.
kind: concept
meaning: An index-side pointer from BadgerDB back to a markdown file + line range
name: FileRef
relationships:
    - target: syde
      type: belongs_to
    - target: storage-engine
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: fileref-7vac
structure_notes: FileRef holds File (relative path), ID, Name, Kind, and Lines (field → [start,end] line range). Stored as the value of the entity index key.
updated_at: "2026-04-14T10:48:03Z"
---
