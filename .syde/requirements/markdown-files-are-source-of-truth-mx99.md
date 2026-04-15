---
id: REQ-0001
kind: requirement
name: Markdown files are source of truth
slug: markdown-files-are-source-of-truth-mx99
description: Entities must persist as markdown + YAML frontmatter so git diffs are human-reviewable.
relationships:
    - target: storage-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:47:01Z"
statement: The syde storage layer shall persist every entity as a markdown file with YAML frontmatter under .syde/<kind-plural>/<slug>.md, and the BadgerDB index shall remain a rebuildable cache that is never authoritative.
req_type: constraint
priority: must
verification: syde reindex round-trips files-to-index without data loss
source: manual
source_ref: decision:DEC-0001
requirement_status: active
rationale: Git-friendly storage lets humans review diffs and resolve merges without special tools. A rebuildable index means corruption is always recoverable via 'syde reindex'.
---
