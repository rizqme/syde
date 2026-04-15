---
id: TSK-0125
kind: task
name: Implement EARS statement validator
slug: implement-ears-statement-validator-d3zu
relationships:
    - target: revamp-requirements-to-ears-coverage-model-dlas
      type: belongs_to
updated_at: "2026-04-15T10:37:39Z"
task_status: completed
priority: high
objective: Non-conforming requirement statements are rejected at save time with a clear error pointing at the allowed patterns.
details: 'internal/model/validation.go: add earsPatternMatch(statement string) returning which of the five patterns matched (Ubiquitous, Event-driven, State-driven, Unwanted-behavior, Optional-feature). Wire into ValidateEntity for RequirementEntity. Bail early if statement is empty.'
acceptance: syde add requirement --statement 'Add a feature' fails with an EARS pattern error; syde add requirement --statement 'The system shall respond within 500ms' succeeds.
affected_entities:
    - entity-model-f28o
    - cli-commands-hpjb
plan_ref: revamp-requirements-to-ears-coverage-model-dlas
plan_phase: phase_2
created_at: "2026-04-15T09:53:47Z"
completed_at: "2026-04-15T10:37:39Z"
---
