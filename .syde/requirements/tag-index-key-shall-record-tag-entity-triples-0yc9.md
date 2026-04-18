---
id: REQ-0329
kind: requirement
name: Tag index key shall record tag-entity triples
slug: tag-index-key-shall-record-tag-entity-triples-0yc9
relationships:
    - target: tag-index-key-5a4q
      type: refines
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:46Z"
statement: The syde storage layer shall store a presence-only entry per (tag, kind, id) triple under BadgerDB key 't:<tag>:<kind>:<id>'.
req_type: interface
priority: must
verification: Integration test tagging an entity and filtering list output by tag
source: manual
source_ref: contract:tag-index-key-5a4q
requirement_status: active
rationale: Tag index keys support O(tag) filtering without walking every entity.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:46Z"
---
