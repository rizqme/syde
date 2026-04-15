---
acceptance: grep -r OpenInMemoryIndex internal/ returns no results; build still passes
completed_at: "2026-04-14T07:00:15Z"
details: Delete the three helpers in internal/storage/index.go + store.go. Verify no callers remain.
id: TSK-0028
kind: task
name: Delete OpenInMemoryIndex / NewInMemoryStore / RebuildIndex
objective: Dead Phase 4 code removed
plan_phase: phase_5
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: delete-openinmemoryindex-newinmemorystore-rebuildindex-3c89
task_status: completed
updated_at: "2026-04-14T07:00:15Z"
---
