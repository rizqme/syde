---
id: FLW-0017
kind: flow
name: Search Model
slug: search-model-5eu6
description: User searches the design model by keyword, code pattern, or file
tags:
    - querying
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-entity-operation-flows
      type: references
updated_at: "2026-04-17T09:12:27Z"
trigger: User runs syde query with search/code/file flags
goal: User finds matching entities with architectural framing
steps:
    - id: s1
      action: User runs syde query --search or --code
      contract: query-entity
      description: CLI dispatches query via HTTP
      on_success: s2
    - id: s2
      action: System searches entities
      contract: search-entities
      description: Matches against entity fields
      on_success: s3
    - id: s3
      action: System returns results via HTTP
      contract: search-entities-http
      description: HTTP API returns annotated hits
      on_success: done
---
