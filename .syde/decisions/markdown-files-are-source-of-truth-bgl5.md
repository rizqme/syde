---
alternatives_considered: SQLite DB (opaque in git diff); pure YAML in one file (no per-entity history); pure BadgerDB (impossible to review in a PR)
category: data
consequences: Every new field must be designed to serialize well in YAML. Reindex must always be idempotent. Skipping the index for a lookup is a bug.
description: Persist entities as git-friendly markdown; treat BadgerDB as a rebuildable cache.
id: DEC-0001
kind: decision
name: Markdown Files Are Source Of Truth
rationale: Git-friendly storage lets humans review diffs and resolve merges without special tools. A rebuildable index means corruption is always recoverable via 'syde reindex'.
relationships:
    - target: syde
      type: applies_to
    - target: storage-engine
      type: applies_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: markdown-files-are-source-of-truth-bgl5
statement: Entities are persisted as markdown + YAML frontmatter under .syde/<kind>/<slug>.md. BadgerDB is a rebuildable index, never authoritative.
tradeoffs: Slower reads than a pure DB; requires serialization round-trips; limited secondary indexing — mitigated by BadgerDB cache
updated_at: "2026-04-14T03:27:02Z"
---
