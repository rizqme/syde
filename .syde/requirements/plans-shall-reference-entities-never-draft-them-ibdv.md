---
id: REQ-0007
kind: requirement
name: Plans shall reference entities, never draft them
slug: plans-shall-reference-entities-never-draft-them-ibdv
description: Plans and phases scope work via --affected-entity and --affected-file; brand-new entities are created inline during implementation.
relationships:
    - target: entity-model-f28o
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T12:48:22Z"
statement: The syde plan model shall not allow plans or phases to embed draft entity definitions.
req_type: constraint
priority: must
verification: audit walking plan entities for embedded drafts
source: manual
source_ref: decision:DEC-0007
requirement_status: superseded
rationale: Draft entities duplicated schema inside plan YAML and led to a fragile 'materialize on execute' step. References are simpler, validator-friendly, and they power the auto-bump of updated_at when a task finishes.
superseded_by:
    - plans-shall-carry-structured-change-diffs-6ah1
---
