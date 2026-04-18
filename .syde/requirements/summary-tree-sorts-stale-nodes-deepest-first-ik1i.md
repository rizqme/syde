---
id: REQ-0192
kind: requirement
name: Summary Tree Sorts Stale Nodes Deepest First
slug: summary-tree-sorts-stale-nodes-deepest-first-ik1i
relationships:
    - target: summary-tree-fq6u
      type: refines
updated_at: "2026-04-18T09:38:05Z"
statement: When StalePaths is requested, the summary tree shall return stale nodes sorted deepest-first so leaves are summarized before their parents.
req_type: functional
priority: must
verification: unit test of StalePaths in summarize.go
source: manual
source_ref: component:summary-tree-fq6u
requirement_status: active
rationale: Deepest-first ordering enforces the leaves-first summarization workflow.
verified_against:
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:38:05Z"
---
