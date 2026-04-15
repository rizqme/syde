---
acceptance: After syde task done <slug> --affected-entity newslug, the task file contains newslug in affected_entities AND newslug's UpdatedAt is bumped
affected_entities:
    - cli-commands
affected_files:
    - internal/cli/task.go
completed_at: "2026-04-14T07:29:45Z"
created_at: "2026-04-14T07:24:22Z"
details: Call store.Update(t, body) after the merge + status flip. touchAffectedEntities reads from t in memory, so the merged set is used automatically. Verify by inspecting the saved markdown.
id: TSK-0033
kind: task
name: Persist merged task before touch cascade
objective: The task file on disk reflects the merged affected set, and touchAffectedEntities bumps the full list
plan_phase: phase_1
plan_ref: task-done-affected-flags-6zl4
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: persist-merged-task-before-touch-cascade-wbm5
task_status: completed
updated_at: "2026-04-14T07:29:45Z"
---
