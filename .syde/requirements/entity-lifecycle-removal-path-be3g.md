---
id: REQ-0024
kind: requirement
name: Entity Lifecycle Removal Path
slug: entity-lifecycle-removal-path-be3g
relationships:
    - target: entity-8x6p
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:46Z"
statement: While an entity exists on disk, the syde CLI shall only remove it via the syde remove command and shall never reuse its ID afterwards.
req_type: functional
priority: must
verification: integration test running syde remove then attempting to observe ID reuse
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Controlled removal keeps history auditable and prevents ID collisions.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:46Z"
---
