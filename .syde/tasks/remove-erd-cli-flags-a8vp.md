---
id: TSK-0079
kind: task
name: Remove ERD CLI flags
slug: remove-erd-cli-flags-a8vp
relationships:
    - target: syde
      type: belongs_to
    - target: clear-all-sync-check-concept-redesign-tasks
      type: references
updated_at: "2026-04-17T09:11:10Z"
task_status: completed
objective: syde add/update concept no longer accepts --attribute, --action, --data-sensitivity, --structure-notes
details: Remove flags and their wiring in add.go and update.go flow cases
acceptance: syde add concept --help shows only --meaning, --invariants, --lifecycle
affected_entities:
    - cli-commands-hpjb
    - cli-commands
affected_files:
    - internal/cli/add.go
    - internal/cli/update.go
plan_ref: concept-entity-redesign-glossary-with-role-based-links
plan_phase: phase_1
created_at: "2026-04-16T11:18:38Z"
completed_at: "2026-04-17T02:52:33Z"
---
