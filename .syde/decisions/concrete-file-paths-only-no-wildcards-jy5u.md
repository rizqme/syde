---
alternatives_considered: Glob patterns for convenience; globs expanded lazily at validation time
category: data
consequences: Every --file flag must pass a literal path that exists in .syde/tree.yaml. Components with many files carry a long --file list. The orphan-file validator can now trust the files list to be complete.
description: Entity files lists must be literal paths — no wildcards or globs.
id: DEC-0006
kind: decision
name: Concrete File Paths Only — No Wildcards
rationale: Every file should be deterministically mappable to its owning component. Wildcards hide drift (new files silently join without review) and defeat the tree-backed orphan-file validator.
relationships:
    - target: syde
      type: applies_to
    - target: entity-model
      type: applies_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: concrete-file-paths-only-no-wildcards-jy5u
statement: Entity 'files' fields must list concrete, literal paths. Globs like 'internal/cli/*.go' are rejected by the validator.
tradeoffs: More typing when creating a component that spans many files. Mitigated by 'syde tree show' to enumerate directory contents cheaply.
updated_at: "2026-04-14T03:27:03Z"
---
