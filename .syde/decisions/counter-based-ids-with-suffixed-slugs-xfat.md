---
alternatives_considered: ULIDs (unreadable in logs); bare slugs without suffix (filename collisions); hash-based IDs (unreadable, unstable under rename)
category: data
consequences: NextID must always pass the live index handle (no fresh-open). Reindex must recompute max counter. 'syde get' must support all three addressing forms.
description: Counter-based IDs (PFX-NNNN) plus 4-char random slug suffixes for uniqueness.
id: DEC-0003
kind: decision
name: Counter-Based IDs With Suffixed Slugs
rationale: Counter IDs are stable, short, and human-readable. Suffixed slugs guarantee filename uniqueness even when two entities share a name. Three addressing forms (full slug, bare slug, parent/child path) cover every lookup scenario while keeping ambiguity explicit.
relationships:
    - target: syde
      type: applies_to
    - target: storage-engine
      type: applies_to
    - target: slug-and-id-utils
      type: applies_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: counter-based-ids-with-suffixed-slugs-xfat
statement: Entity IDs are counter-based per kind (SYS-0001, COM-0001, CON-0001, ...). Slugs are a name-slugified base plus a 4-char random alphanumeric suffix. Counters live in BadgerDB and are recalculated during reindex from the highest observed ID.
tradeoffs: Counters require transactional allocation (BadgerDB lock). Bare-slug lookups can be ambiguous — surfaced as an error with the candidate list.
updated_at: "2026-04-14T03:27:03Z"
---
