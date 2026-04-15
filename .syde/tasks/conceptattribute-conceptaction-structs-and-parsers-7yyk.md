---
acceptance: 'Unit smoke: syde add concept ''Smoke'' --meaning foo --attribute ''id|uuid|primary'' round-trips through FileStore.Load and returns the same attribute slice.'
affected_entities:
    - entity-model
affected_files:
    - internal/model/entity.go
completed_at: "2026-04-14T09:46:28Z"
created_at: "2026-04-14T09:44:26Z"
details: 'internal/model/entity.go: (1) type ConceptAttribute struct {Name, Type, Description string} with yaml tags. (2) type ConceptAction struct {Name, Description string}. (3) ParseConceptAttribute(spec string) (ConceptAttribute, bool) and ParseConceptAction(spec string) (ConceptAction, bool) mirroring ParseContractParam. (4) Extend ConceptEntity with Attributes []ConceptAttribute and Actions []ConceptAction, yaml omitempty.'
id: TSK-0057
kind: task
name: ConceptAttribute + ConceptAction structs and parsers
objective: ConceptEntity YAML round-trips Attributes and Actions slices; pipe parsers accept the documented spec
plan_phase: phase_1
plan_ref: concept-as-erd
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: conceptattribute-conceptaction-structs-and-parsers-7yyk
task_status: completed
updated_at: "2026-04-14T09:46:28Z"
---
