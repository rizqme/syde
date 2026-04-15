---
acceptance: sync validation reports malformed requirement statuses, missing superseded_by links, and obsolete requirements without obsolete_reason.
affected_entities:
    - audit-engine-4ktg
    - cli-commands-hpjb
affected_files:
    - internal/audit
    - internal/cli/update.go
    - internal/model/entity.go
    - internal/audit/requirements.go
    - internal/audit/audit.go
    - internal/model/validation.go
completed_at: "2026-04-15T06:32:57Z"
created_at: "2026-04-15T06:31:23Z"
details: Extend audit validation so superseded requirements point at replacing requirements, obsolete requirements carry reasons, and requirement status values stay constrained. Use the update command's existing supersedes/obsolete flags as the command flow.
id: TSK-0095
kind: task
name: Validate requirement conflict lifecycle
objective: Validate requirement supersede and obsolete status rules without deleting historical requirements.
plan_phase: phase_4
plan_ref: add-requirement-entity-56dv
priority: high
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: validate-requirement-conflict-lifecycle-txw9
task_status: completed
updated_at: "2026-04-15T06:32:57Z"
---
