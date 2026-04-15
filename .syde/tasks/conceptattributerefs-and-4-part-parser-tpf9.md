---
acceptance: syde add concept X --meaning m --attribute 'a|string|desc' --attribute 'b|uuid|FK|other' round-trips through YAML and the second attribute has refs=[other].
affected_entities:
    - entity-model
affected_files:
    - internal/model/entity.go
completed_at: "2026-04-14T10:26:48Z"
created_at: "2026-04-14T10:22:14Z"
details: 'internal/model/entity.go: ConceptAttribute gains Refs []string with yaml/json tags refs,omitempty. ParseConceptAttribute changes SplitN from 3 to 4; when parts[3] is present, split on '','' with TrimSpace, filter empty. skill docs updated to show the 4-part example (customer_id|uuid|foreign key|customer).'
id: TSK-0067
kind: task
name: ConceptAttribute.Refs and 4-part parser
objective: Attribute spec accepts a 4th pipe field with comma-separated concept slugs the attribute references
plan_phase: phase_5
plan_ref: erd-inside-concept-view
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: conceptattributerefs-and-4-part-parser-tpf9
task_status: completed
updated_at: "2026-04-14T10:26:48Z"
---
