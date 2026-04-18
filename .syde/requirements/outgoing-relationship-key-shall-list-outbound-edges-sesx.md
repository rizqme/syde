---
id: REQ-0327
kind: requirement
name: Outgoing relationship key shall list outbound edges
slug: outgoing-relationship-key-shall-list-outbound-edges-sesx
relationships:
    - target: outgoing-relationship-index-key-45p4
      type: refines
    - target: storage-engine-ahgm
      type: refines
updated_at: "2026-04-18T09:36:47Z"
statement: The syde storage layer shall store the set of outbound relationships from a source entity under BadgerDB key 'r:out:<id>'.
req_type: interface
priority: must
verification: Integration test adding a relationship and reading r:out:<id>
source: manual
source_ref: contract:outgoing-relationship-index-key-45p4
requirement_status: active
rationale: Outbound relationship keys support fast traversal for graph views and traceability audits.
verified_against:
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:36:47Z"
---
