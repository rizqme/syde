---
id: REQ-0006
kind: requirement
name: Entity file lists shall contain only literal paths
slug: entity-file-lists-shall-contain-only-literal-paths-hatq
description: No wildcards or globs in a file list.
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:46:50Z"
statement: The syde validator shall reject any files entry that contains a wildcard or that does not exist in the summary tree.
req_type: constraint
priority: must
verification: automated test invoking the validator with a wildcard path
source: manual
source_ref: decision:DEC-0006
requirement_status: active
rationale: Every file should be deterministically mappable to its owning component. Wildcards hide drift (new files silently join without review) and defeat the tree-backed orphan-file validator.
---
