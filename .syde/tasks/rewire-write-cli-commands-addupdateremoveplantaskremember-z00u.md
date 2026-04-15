---
acceptance: syde add component creates the file AND updates the index (visible via syde query immediately)
completed_at: "2026-04-14T06:58:09Z"
details: Keep FileStore usage local. After FileStore.Save, call client.Reindex([]string{relPath}). Remove direct s.Idx.IndexEntity calls from write paths — that's now syded's job. Apply to add, update, remove, plan, task, remember, design, open.
id: TSK-0024
kind: task
name: Rewire write CLI commands (add/update/remove/plan/task/remember)
objective: Write commands save markdown via FileStore then call client.Reindex
plan_phase: phase_4
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: rewire-write-cli-commands-addupdateremoveplantaskremember-z00u
task_status: completed
updated_at: "2026-04-14T06:58:09Z"
---
