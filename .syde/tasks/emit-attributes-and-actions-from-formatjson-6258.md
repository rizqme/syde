---
acceptance: curl .../api/<proj>/entity/decision-m2um | jq '.entity.attributes' returns a non-empty list after rebuild and daemon restart.
affected_entities:
    - query-engine
affected_files:
    - internal/query/formatter.go
completed_at: "2026-04-14T10:23:15Z"
created_at: "2026-04-14T10:22:14Z"
details: internal/query/formatter.go FormatJSON case *model.ConceptEntity adds entityMap['attributes'] = e.Attributes, entityMap['actions'] = e.Actions, entityMap['concept_relationships'] = e.ConceptRelationships.
id: TSK-0063
kind: task
name: Emit attributes and actions from FormatJSON
objective: GET /api/<proj>/entity/<slug> returns attributes + actions in the JSON body for concept entities
plan_phase: phase_1
plan_ref: erd-inside-concept-view
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: emit-attributes-and-actions-from-formatjson-6258
task_status: completed
updated_at: "2026-04-14T10:23:15Z"
---
