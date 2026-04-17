---
id: TSK-0001
kind: task
name: Add PlanEntity Design field and Changes struct
slug: add-planentity-design-field-and-changes-struct-r602
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-carry-structured-change-diffs-6ah1
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: internal/model/plan.go defines Design string + Changes PlanChanges and PlanChanges groups per-kind ChangeLanes.
details: Add Design string YAML-tagged 'design,omitempty' to PlanEntity. Add PlanChanges with six ChangeLane fields (Requirements/Systems/Concepts/Components/Contracts/Flows) YAML-tagged snake_case. ChangeLane has Deleted []DeletedChange, Extended []ExtendedChange, New []NewChange. DeletedChange {ID, Slug, Why}. ExtendedChange {ID, Slug, What, Why, FieldChanges map[string]string}. NewChange {ID, Name, What, Why, Draft RawNewDraft} where RawNewDraft is a map[string]interface{} for kind-specific fields. IDs are 4-char ULID-style tokens generated on insertion.
acceptance: go build ./... compiles, unmarshalling a YAML plan with design and changes round-trips correctly.
affected_entities:
    - entity-model-f28o
affected_files:
    - internal/model/plan.go
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_1
created_at: "2026-04-15T11:40:36Z"
completed_at: "2026-04-15T11:46:57Z"
---
