---
id: REQ-0125
kind: requirement
name: Tree Node Path Unique In Tree
slug: tree-node-path-unique-in-tree-zgei
relationships:
    - target: tree-node-iutv
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:49Z"
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
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:49Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:36:49Z"
---
