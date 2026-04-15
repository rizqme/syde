---
id: TSK-0123
kind: task
name: Add req_type priority verification fields
slug: add-reqtype-priority-verification-fields-5nvi
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:34:20Z"
task_status: completed
priority: high
objective: RequirementEntity has req_type and priority enum fields plus a verification field replacing acceptance_criteria, with YAML tags and validation.
details: 'internal/model/entity.go: add RequirementType enum (functional, non-functional, constraint, interface, performance, security, usability); add RequirementPriority enum (must, should, could, wont); add Verification string field; remove or alias AcceptanceCriteria. Update validation.go to enforce enum membership when fields are set.'
acceptance: go build ./... succeeds; syde add requirement --type functional --priority must --verification 'automated test' creates a requirement with those fields.
affected_entities:
    - entity-model-f28o
    - cli-commands-hpjb
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_2
created_at: "2026-04-15T09:53:47Z"
completed_at: "2026-04-15T10:34:20Z"
---
