---
acceptance: A unique word appearing only in a component body is returned by ./syde query --search <word>.
affected_entities:
    - storage-engine
affected_files:
    - internal/storage/indexer.go
completed_at: "2026-04-14T08:14:37Z"
created_at: "2026-04-14T08:07:21Z"
details: In internal/storage/indexer.go where words are collected, for the six substantive kinds also tokenize the body content (below frontmatter). Skip plan/task/system/subsystem whose bodies are procedural/churny.
id: TSK-0042
kind: task
name: Tokenize entity body markdown in the indexer
objective: Search matches terms in component/contract/concept/decision/flow/learning bodies, not just frontmatter fields
plan_phase: phase_1
plan_ref: improve-syde-query
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: tokenize-entity-body-markdown-in-the-indexer-x99x
task_status: completed
updated_at: "2026-04-14T08:14:37Z"
---
