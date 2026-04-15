---
acceptance: After reindex, syde query --search ConceptEntity finds the Entity concept (or any entity whose body mentions ConceptEntity), and --search IndexSchemaVersion finds the storage-engine learning written this session.
affected_entities:
    - storage-engine
affected_files:
    - internal/storage/indexer.go
    - internal/storage/index.go
completed_at: "2026-04-14T09:28:36Z"
created_at: "2026-04-14T09:20:18Z"
details: 'internal/storage/indexer.go: extend tokenize() to also split CamelCase boundaries (lowercase->uppercase transitions) and split on _ and -. Emit both the original concatenated form and each sub-token. Filter stop words and len<2 as today. internal/storage/index.go bump IndexSchemaVersion = 3. NewStore auto-reindex picks it up on next open.'
id: TSK-0054
kind: task
name: CamelCase / snake_case tokenizer split with IndexSchemaVersion v3
objective: Tokens like ConceptEntity, IndexSchemaVersion, snake_case_field, dash-cased-name index every reasonable sub-form so loose searches still hit
plan_phase: phase_2
plan_ref: syde-context-bridge
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: camelcase-snakecase-tokenizer-split-with-indexschemaversion-v3-v5c9
task_status: completed
updated_at: "2026-04-14T09:28:36Z"
---
