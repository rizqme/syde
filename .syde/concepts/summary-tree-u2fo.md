---
attributes:
    - description: last scan wall clock
      name: scanned_at
    - description: project root path, always '.'
      name: root
    - description: flat path-keyed node map
      name: nodes
description: Mirror of the source tree with human-written summaries and cascade-stale tracking.
id: CPT-0010
invariants: Every tracked file must have a hash. A file changing hash marks itself + all ancestors stale. A folder summary is marked stale when any child's summary changes.
kind: concept
lifecycle: scan → diff → mark stale → summarize → cascade stale up to root → clean
meaning: A mirror of the source tree where every file and folder carries a human-written summary plus change tracking
name: Summary Tree
relationships:
    - target: syde
      type: belongs_to
    - target: summary-tree
      type: references
    - target: tree-node
      type: relates_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
slug: summary-tree-u2fo
structure_notes: Persisted as .syde/tree.yaml. Flat map keyed by relative path. Contains metadata scanned_at, root, and nodes.
updated_at: "2026-04-14T10:48:03Z"
---
