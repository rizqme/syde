---
id: CPT-0010
kind: concept
name: Summary Tree
slug: summary-tree-u2fo
description: Mirror of the source tree with human-written summaries and cascade-stale tracking.
relationships:
    - target: syde
      type: belongs_to
    - target: tree-node
      type: relates_to
    - target: summary-tree-fq6u
      type: implemented_by
updated_at: "2026-04-17T08:25:52Z"
meaning: A mirror of the source tree where every file and folder carries a human-written summary plus change tracking
lifecycle: scan → diff → mark stale → summarize → cascade stale up to root → clean
invariants: Every tracked file must have a hash. A file changing hash marks itself + all ancestors stale. A folder summary is marked stale when any child's summary changes.
---
