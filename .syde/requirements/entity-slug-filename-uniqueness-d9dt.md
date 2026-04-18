---
id: REQ-0022
kind: requirement
name: Entity Slug Filename Uniqueness
slug: entity-slug-filename-uniqueness-d9dt
relationships:
    - target: entity-8x6p
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:05Z"
statement: The syde CLI shall ensure that every entity slug filename is unique on disk under .syde/.
req_type: constraint
priority: must
verification: integration test attempting duplicate slug creation
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: Slug collisions would overwrite entity files silently.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:05Z"
---
