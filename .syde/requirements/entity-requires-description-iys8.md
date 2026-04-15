---
id: REQ-0031
kind: requirement
name: Entity Requires Description
slug: entity-requires-description-iys8
relationships:
    - target: entity-8x6p
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:04Z"
statement: The syde CLI shall require a non-empty description on every entity instance.
req_type: constraint
priority: must
verification: integration test running syde add <kind> without --description
source: manual
source_ref: concept:entity-8x6p
requirement_status: active
rationale: The description is the one-sentence elevator pitch used by list and search commands.
---
