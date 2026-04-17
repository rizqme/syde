---
id: FLW-0033
kind: flow
name: Reindex Storage
slug: reindex-storage-ekiq
description: User rebuilds the BadgerDB index from markdown files
tags:
    - init
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User runs syde reindex or index is corrupted
goal: BadgerDB index matches the current markdown files
steps:
    - id: s1
      action: User runs syde reindex
      contract: reindex-from-files
      description: Rebuilds all BadgerDB keys from .syde/ files
      on_success: done
---
