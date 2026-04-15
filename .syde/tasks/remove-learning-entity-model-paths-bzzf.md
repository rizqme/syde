---
id: TSK-0112
kind: task
name: Remove learning entity model paths
slug: remove-learning-entity-model-paths-bzzf
relationships:
    - target: remove-learning-entity-altogether-wk7c
      type: references
    - target: limit-baseline-style-requirement-fanout-m156
      type: references
    - target: remove-learning-entity-and-cap-requirement-fanout-wqyz
      type: belongs_to
updated_at: "2026-04-15T09:23:33Z"
task_status: completed
priority: high
objective: Remove KindLearning, LearningEntity, and learning-specific query/memory model hooks.
details: Edit model entity kind dispatch, validation, learning model file, query resolver/engine counters, and memory manager so learning is no longer a first-class entity.
acceptance: rg finds no KindLearning or LearningEntity references in Go code.
affected_entities:
    - entity-model-f28o
    - query-engine-9k84
    - memory-sync-hgir
affected_files:
    - internal/model/entity.go
    - internal/model/learning.go
    - internal/model/validation.go
    - internal/query/resolver.go
    - internal/query/engine.go
    - internal/memory/manager.go
plan_ref: remove-learning-entity-and-cap-requirement-fanout-wqyz
plan_phase: phase_1
created_at: "2026-04-15T08:23:05Z"
completed_at: "2026-04-15T09:00:53Z"
---
