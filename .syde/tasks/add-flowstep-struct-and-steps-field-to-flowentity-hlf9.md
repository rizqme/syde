---
id: TSK-0056
kind: task
name: Add FlowStep struct and Steps field to FlowEntity
slug: add-flowstep-struct-and-steps-field-to-flowentity-hlf9
relationships:
    - target: syde
      type: belongs_to
    - target: flow-steps-flow-authoring-tasks
      type: references
updated_at: "2026-04-17T09:14:43Z"
task_status: completed
objective: FlowEntity.Steps []FlowStep exists with proper yaml/json tags
details: Add FlowStep struct to internal/model/entity.go with Action, Contract, Description, OnSuccess, OnFailure string fields. Add Steps []FlowStep yaml:steps,omitempty json:steps,omitempty to FlowEntity.
acceptance: go build clean; round-trip YAML with steps preserves the list
affected_entities:
    - entity-model-f28o
affected_files:
    - internal/model/entity.go
plan_ref: flow-steps-with-contract-references-and-flowchart-rendering
plan_phase: phase_1
created_at: "2026-04-16T09:22:11Z"
completed_at: "2026-04-16T10:32:06Z"
---
