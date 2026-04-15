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
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T11:03:44Z"
statement: The syde storage layer shall store a presence-only entry per (tag, kind, id) triple under BadgerDB key 't:<tag>:<kind>:<id>'.
req_type: interface
priority: must
verification: Integration test tagging an entity and filtering list output by tag
source: manual
source_ref: contract:tag-index-key-5a4q
requirement_status: active
rationale: Tag index keys support O(tag) filtering without walking every entity.
---
