---
id: CPT-0005
kind: concept
name: Tree Node
slug: tree-node-iutv
description: One file or folder entry in the source-mirror summary tree.
relationships:
    - target: syde
      type: belongs_to
    - target: summary-tree-fq6u
      type: implemented_by
updated_at: "2026-04-17T08:25:52Z"
meaning: A single file or folder entry in the summary tree mirroring the source tree
lifecycle: discovered by walk → hashed → added → summarized → stale on content change → re-summarized → removed when file deleted
invariants: Path is unique within the tree. Parent must exist (except for root). Hash must match file bytes when scan runs. Root (.) has no parent and is never stale via cascade alone.
---
