---
id: REQ-0022
kind: requirement
name: Entity Slug Filename Uniqueness
slug: entity-slug-filename-uniqueness-d9dt
relationships:
    - target: entity-8x6p
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:51:54Z"
statement: The syde CLI shall ensure that every entity slug filename is unique on disk under .syde/.
req_type: constraint
priority: must
verification: integration test attempting duplicate slug creation
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Slug collisions would overwrite entity files silently.
---
