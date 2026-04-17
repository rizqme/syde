---
id: REQ-0125
kind: requirement
name: Tree Node Path Unique In Tree
slug: tree-node-path-unique-in-tree-zgei
relationships:
    - target: tree-node-iutv
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:46:07Z"
statement: The syde CLI shall ensure that every tree node path is unique within the summary tree.
req_type: constraint
priority: must
verification: unit test loading a tree.yaml with duplicate paths
source: manual
source_ref: concept:tree-node-iutv
requirement_status: active
rationale: The flat path-keyed map cannot tolerate duplicate keys without data loss.
audited_overlaps:
    - slug: plan-phase-id-unique-within-plan-1wq7
      distinction: Tree node path uniqueness scopes to the summary tree; phase ID uniqueness scopes within a parent plan, different entity spaces and identifier types.
---
