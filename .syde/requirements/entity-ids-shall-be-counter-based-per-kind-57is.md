---
id: REQ-0003
kind: requirement
name: Entity IDs shall be counter-based per kind
slug: entity-ids-shall-be-counter-based-per-kind-57is
description: IDs are short human-readable counter form plus a slug suffix for filename uniqueness.
relationships:
    - target: storage-engine
      type: refines
    - target: slug-and-id-utils
      type: refines
updated_at: "2026-04-18T09:38:00Z"
statement: The syde storage layer shall allocate entity IDs as a counter-based prefix per kind (SYS-0001, COM-0001, CON-0001, ...) and shall assign each entity a slug formed as the name-slugified base plus a 4-character random alphanumeric suffix.
req_type: constraint
priority: must
verification: audit verifying every entity ID matches the per-kind counter prefix
source: manual
source_ref: decision:DEC-0003
requirement_status: active
rationale: Counter IDs are stable, short, and human-readable. Suffixed slugs guarantee filename uniqueness even when two entities share a name.
verified_against:
    slug-and-id-utils-8kr7:
        hash: 2a28c2d9c9e40b4ca1b47bbbf49b2face3e0b4599f68eb1f6c0520d4258c3d4c
        at: "2026-04-18T09:38:00Z"
    storage-engine-ahgm:
        hash: f360017cda1e57fe0083d2f867db63e847625a33a670b76215d7787f434555c3
        at: "2026-04-18T09:38:00Z"
---
