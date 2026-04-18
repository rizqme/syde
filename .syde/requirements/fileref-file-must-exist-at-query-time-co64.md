---
id: REQ-0035
kind: requirement
name: FileRef File Must Exist At Query Time
slug: fileref-file-must-exist-at-query-time-co64
relationships:
    - target: fileref-7vac
      type: refines
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:56Z"
statement: If a FileRef points to a markdown file that no longer exists on disk, then the storage layer shall report a stale-index error and trigger a reindex.
req_type: constraint
priority: must
verification: integration test deleting a file then querying the index
source: manual
source_ref: concept:fileref-7vac
requirement_status: active
rationale: Stale file refs break navigation from index results back to source.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:56Z"
---
