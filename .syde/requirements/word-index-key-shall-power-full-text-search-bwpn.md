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
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T11:03:44Z"
statement: The syde storage layer shall store a tokenized word-to-entity field mapping under BadgerDB key 'w:<word>:<kind>:<id>'.
req_type: interface
priority: must
verification: Integration test indexing a markdown body and searching for a tokenized word
source: manual
source_ref: contract:word-index-key-4s3r
requirement_status: active
rationale: Word index keys back the CLI and dashboard full-text search path.
---
