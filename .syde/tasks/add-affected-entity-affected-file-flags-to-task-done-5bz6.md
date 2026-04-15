---
acceptance: syde task done <slug> --affected-entity foo --affected-file bar parses; --help lists them
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/task.go
completed_at: "2026-04-14T07:29:45Z"
created_at: "2026-04-14T07:24:22Z"
details: Register taskDoneAffectedEntities / taskDoneAffectedFiles as StringArrayVar flags on taskDoneCmd. Repeatable, same syntax as create.
id: TSK-0031
kind: task
name: Add --affected-entity / --affected-file flags to task done
objective: CLI grammar accepts the new flags on task done
plan_phase: phase_1
plan_ref: task-done-affected-flags-6zl4
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-affected-entity-affected-file-flags-to-task-done-5bz6
task_status: completed
updated_at: "2026-04-14T07:29:45Z"
---
