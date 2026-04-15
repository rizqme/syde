---
id: REQ-0157
kind: requirement
name: Query Engine Returns Git Diff For Entity
slug: query-engine-returns-git-diff-for-entity-ywgr
relationships:
    - target: query-engine-9k84
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:10Z"
statement: When EntityDiff is called for an entity, the query engine shall return recent change entries from git log for the entity markdown file.
req_type: functional
priority: should
verification: unit test of EntityDiff in internal/query/diff.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Git-sourced history is the audit log for entity evolution.
---
