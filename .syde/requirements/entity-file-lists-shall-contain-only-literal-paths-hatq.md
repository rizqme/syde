---
id: REQ-0006
kind: requirement
name: Entity file lists shall contain only literal paths
slug: entity-file-lists-shall-contain-only-literal-paths-hatq
description: No wildcards or globs in a file list.
relationships:
    - target: entity-model-f28o
      type: refines
updated_at: "2026-04-18T09:36:44Z"
statement: The syde validator shall reject any files entry that contains a wildcard or that does not exist in the summary tree.
req_type: constraint
priority: must
verification: automated test invoking the validator with a wildcard path
source: manual
source_ref: decision:DEC-0006
requirement_status: active
rationale: Every file should be deterministically mappable to its owning component. Wildcards hide drift (new files silently join without review) and defeat the tree-backed orphan-file validator.
verified_against:
    entity-model-f28o:
        hash: 7e51689e4dc181c602291eabd785a2d15d5fe4750220e6782ab3d61c0640b0b8
        at: "2026-04-18T09:36:44Z"
---
