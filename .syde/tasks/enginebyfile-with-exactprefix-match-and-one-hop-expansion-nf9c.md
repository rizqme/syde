---
acceptance: ./syde query --file internal/storage/index.go lists storage-engine plus query-engine (inbound depends_on) and related items.
affected_entities:
    - query-engine
affected_files:
    - internal/query/engine.go
    - internal/query/resolver.go
completed_at: "2026-04-14T08:18:48Z"
created_at: "2026-04-14T08:07:36Z"
details: ByFileResult {Path, Owners []EntitySummary, Related []EntitySummary}. Walk storage file coverage for owners. If path ends with / or has no exact match, prefix-filter tree node paths. If withRelated, for each owner add outbound rels + inbound rels resolved via Index, de-duped by ID.
id: TSK-0044
kind: task
name: Engine.ByFile with exact/prefix match and one-hop expansion
objective: Given a source file or directory prefix, return owning entities plus their one-hop neighbors
plan_phase: phase_2
plan_ref: improve-syde-query
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: enginebyfile-with-exactprefix-match-and-one-hop-expansion-nf9c
task_status: completed
updated_at: "2026-04-14T08:18:48Z"
---
