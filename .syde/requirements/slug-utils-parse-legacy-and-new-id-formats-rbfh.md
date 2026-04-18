---
id: REQ-0174
kind: requirement
name: Slug Utils Parse Legacy And New ID Formats
slug: slug-utils-parse-legacy-and-new-id-formats-rbfh
relationships:
    - target: slug-and-id-utils-8kr7
      type: refines
updated_at: "2026-04-18T09:38:05Z"
statement: When ParseIDKind is called with an ID string, the slug utils shall extract the EntityKind from either the new PFX-NNNN or legacy pfx_xxxxxxxx format.
req_type: functional
priority: must
verification: unit test of ParseIDKind in internal/utils/id.go
source: manual
source_ref: component:slug-and-id-utils-8kr7
requirement_status: active
rationale: Legacy IDs still appear in historical files and must remain resolvable.
verified_against:
    slug-and-id-utils-8kr7:
        hash: 2a28c2d9c9e40b4ca1b47bbbf49b2face3e0b4599f68eb1f6c0520d4258c3d4c
        at: "2026-04-18T09:38:05Z"
---
