---
acceptance: Passing an unknown --affected-entity slug errors out and leaves the task unchanged
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/task.go
completed_at: "2026-04-14T07:29:45Z"
created_at: "2026-04-14T07:24:22Z"
details: 'Inside taskDoneCmd.RunE, after loading the task via store.Get: append flag values; dedupe via simple set; call validateTaskReferences on the additions (the pre-existing values are already valid). Abort on validation error.'
id: TSK-0032
kind: task
name: Merge flag values into the loaded task + validate
objective: Loaded task has the full affected set before status is flipped
plan_phase: phase_1
plan_ref: task-done-affected-flags-6zl4
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: merge-flag-values-into-the-loaded-task-validate-1amv
task_status: completed
updated_at: "2026-04-14T07:29:45Z"
---
