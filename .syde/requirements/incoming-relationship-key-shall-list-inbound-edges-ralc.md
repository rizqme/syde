---
id: REQ-0326
kind: requirement
name: Incoming relationship key shall list inbound edges
slug: incoming-relationship-key-shall-list-inbound-edges-ralc
relationships:
    - target: incoming-relationship-index-key-lro0
      type: refines
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:37:53Z"
statement: The syde storage layer shall store the set of inbound relationship sources for a target entity under BadgerDB key 'r:in:<target>'.
req_type: interface
priority: must
verification: Integration test adding a relationship and reading r:in:<target>
source: manual
source_ref: contract:incoming-relationship-index-key-lro0
requirement_status: active
rationale: Inbound relationship keys let impact analysis walk upstream in O(1) per hop.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:37:53Z"
---
