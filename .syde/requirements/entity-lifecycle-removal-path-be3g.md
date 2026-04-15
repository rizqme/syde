---
id: REQ-0024
kind: requirement
name: Entity Lifecycle Removal Path
slug: entity-lifecycle-removal-path-be3g
relationships:
    - target: entity-8x6p
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:51:59Z"
statement: While an entity exists on disk, the syde CLI shall only remove it via the syde remove command and shall never reuse its ID afterwards.
req_type: functional
priority: must
verification: integration test running syde remove then attempting to observe ID reuse
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Controlled removal keeps history auditable and prevents ID collisions.
---
