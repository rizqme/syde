---
id: REQ-0011
kind: requirement
name: Markdown shall be the source of truth
slug: markdown-shall-be-the-source-of-truth-urm4
relationships:
    - target: syde-5tdt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:49:02Z"
statement: The syde storage layer shall treat markdown files under .syde as the only authoritative entity representation.
req_type: constraint
priority: must
verification: syde reindex rebuilds the BadgerDB index from files alone with no data loss
source: manual
source_ref: system:syde-5tdt:design_principles
requirement_status: active
rationale: Git-friendly text storage lets humans diff and merge without special tools.
---
