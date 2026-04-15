---
alternatives_considered: Keep draft entities; add a validation layer on drafts
category: process
consequences: Removed plan add-entity command. Task create requires --affected-entity/--affected-file. 'task done' auto-bumps updated_at on every listed entity plus every entity whose files overlap affected_files.
description: Plans no longer carry draft entities; tasks reference existing entities and files.
id: DEC-0007
kind: decision
name: Plans Reference Entities, Never Draft Them
rationale: Draft entities duplicated schema inside plan YAML and led to a fragile 'materialize on execute' step. References are simpler, validator-friendly, and they power the auto-bump of updated_at when a task finishes — directly clearing drift warnings.
relationships:
    - target: syde
      type: applies_to
    - target: plan
      type: applies_to
    - target: task
      type: applies_to
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: plans-reference-entities-never-draft-them-thvu
statement: Plans and phases do not carry draft entities. Tasks declare their scope via --affected-entity (existing slugs) and --affected-file (existing tree paths). Brand-new entities are created inline during implementation via 'syde add'.
tradeoffs: Plans expressing 'we will create a new component' must describe the intent in phase --notes rather than a typed draft.
updated_at: "2026-04-14T03:27:03Z"
---
