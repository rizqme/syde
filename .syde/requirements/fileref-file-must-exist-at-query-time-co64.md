---
id: REQ-0035
kind: requirement
name: FileRef File Must Exist At Query Time
slug: fileref-file-must-exist-at-query-time-co64
relationships:
    - target: fileref-7vac
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:14Z"
statement: If a FileRef points to a markdown file that no longer exists on disk, then the storage layer shall report a stale-index error and trigger a reindex.
req_type: constraint
priority: must
verification: integration test deleting a file then querying the index
source: manual
source_ref: concept:fileref-7vac
requirement_status: active
rationale: Stale file refs break navigation from index results back to source.
---
