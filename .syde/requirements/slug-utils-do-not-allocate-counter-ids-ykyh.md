---
id: REQ-0175
kind: requirement
name: Slug Utils Do Not Allocate Counter IDs
slug: slug-utils-do-not-allocate-counter-ids-ykyh
relationships:
    - target: slug-and-id-utils-8kr7
      type: refines
updated_at: "2026-04-18T09:37:23Z"
statement: The slug utils shall not allocate counter-based entity IDs and shall defer allocation to the storage engine counter module.
req_type: constraint
priority: must
verification: code review of internal/utils for counter writes
source: manual
source_ref: component:slug-and-id-utils-8kr7
requirement_status: active
rationale: Counter allocation requires BadgerDB access that is owned by storage.
verified_against:
    slug-and-id-utils-8kr7:
        hash: 2a28c2d9c9e40b4ca1b47bbbf49b2face3e0b4599f68eb1f6c0520d4258c3d4c
        at: "2026-04-18T09:37:23Z"
---
