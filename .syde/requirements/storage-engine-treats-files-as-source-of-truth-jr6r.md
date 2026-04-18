---
id: REQ-0184
kind: requirement
name: Storage Engine Treats Files As Source Of Truth
slug: storage-engine-treats-files-as-source-of-truth-jr6r
relationships:
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:36:53Z"
statement: The storage engine shall treat markdown files as the source of truth and shall allow the BadgerDB index to be rebuilt from them at any time.
req_type: constraint
priority: must
verification: integration test that deleting index and reindexing restores state
source: manual
source_ref: component:storage-engine-ahgm
requirement_status: active
rationale: Git-friendly text files must outrank any derived cache.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:36:53Z"
---
