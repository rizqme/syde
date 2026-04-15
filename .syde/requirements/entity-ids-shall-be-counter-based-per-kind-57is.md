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
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:47:01Z"
statement: The syde storage layer shall allocate entity IDs as a counter-based prefix per kind (SYS-0001, COM-0001, CON-0001, ...) and shall assign each entity a slug formed as the name-slugified base plus a 4-character random alphanumeric suffix.
req_type: constraint
priority: must
verification: audit verifying every entity ID matches the per-kind counter prefix
source: manual
source_ref: decision:DEC-0003
requirement_status: active
rationale: Counter IDs are stable, short, and human-readable. Suffixed slugs guarantee filename uniqueness even when two entities share a name.
---
