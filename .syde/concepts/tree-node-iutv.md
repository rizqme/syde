---
attributes:
    - description: relative path from project root
      name: path
    - description: file or dir
      name: type
    - description: parent directory path
      name: parent
    - description: direct child paths (dirs only)
      name: children
    - description: file size in bytes (files only)
      name: size
    - description: sha256 of file content (files only)
      name: hash
    - description: filesystem mtime (files only)
      name: mtime
    - description: set when sniffer detected binary content
      name: binary
    - description: human-written 1-3 sentence description
      name: summary
    - description: set when file changed after last summarize
      name: summary_stale
    - description: when true, exempt from orphan/stale gates
      name: ignored
description: One file or folder entry in the source-mirror summary tree.
id: CPT-0005
invariants: Path is unique within the tree. Parent must exist (except for root). Hash must match file bytes when scan runs. Root (.) has no parent and is never stale via cascade alone.
kind: concept
lifecycle: discovered by walk → hashed → added → summarized → stale on content change → re-summarized → removed when file deleted
meaning: A single file or folder entry in the summary tree mirroring the source tree
name: Tree Node
relationships:
    - target: syde
      type: belongs_to
    - target: summary-tree
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: tree-node-iutv
structure_notes: Node has type (file|dir), parent path, children list, size, hash (SHA-256), mtime, binary flag, summary text, summary_stale flag, updated_at, ignored flag. Stored in the flat Nodes map of .syde/tree.yaml.
updated_at: "2026-04-14T10:48:03Z"
---
