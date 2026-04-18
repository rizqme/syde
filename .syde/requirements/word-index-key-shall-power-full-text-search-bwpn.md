---
id: REQ-0330
kind: requirement
name: Word index key shall power full-text search
slug: word-index-key-shall-power-full-text-search-bwpn
relationships:
    - target: word-index-key-4s3r
      type: refines
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:36:54Z"
statement: The syde storage layer shall store a tokenized word-to-entity field mapping under BadgerDB key 'w:<word>:<kind>:<id>'.
req_type: interface
priority: must
verification: Integration test indexing a markdown body and searching for a tokenized word
source: manual
source_ref: contract:word-index-key-4s3r
requirement_status: active
rationale: Word index keys back the CLI and dashboard full-text search path.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:36:54Z"
---
