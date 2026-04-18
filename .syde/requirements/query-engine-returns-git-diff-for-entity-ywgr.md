---
id: REQ-0157
kind: requirement
name: Query Engine Returns Git Diff For Entity
slug: query-engine-returns-git-diff-for-entity-ywgr
relationships:
    - target: query-engine-9k84
      type: refines
updated_at: "2026-04-18T09:37:35Z"
statement: When EntityDiff is called for an entity, the query engine shall return recent change entries from git log for the entity markdown file.
req_type: functional
priority: should
verification: unit test of EntityDiff in internal/query/diff.go
source: manual
source_ref: component:query-engine-9k84
requirement_status: active
rationale: Git-sourced history is the audit log for entity evolution.
verified_against:
    query-engine-9k84:
        hash: 03a24974e906ccbc86ac65d8d2da018434bef5290e59b82647d94ff0290ac1d3
        at: "2026-04-18T09:37:35Z"
---
