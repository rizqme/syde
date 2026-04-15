---
acceptance: 'Creating a new component under an existing system bumps the system''s UpdatedAt; deleting it also bumps the system; chain: creating a contract under a sub-system bumps the sub-system AND the root system above it'
affected_entities:
    - storage-engine
affected_files:
    - internal/storage/store.go
completed_at: "2026-04-14T07:31:57Z"
created_at: "2026-04-14T07:26:27Z"
details: Add s.CreateCascade(e, body) which wraps s.Create and then calls cascadeFromParent on the new entity's belongs_to targets. Add s.DeleteCascade(kind, slug) which loads the entity first (to capture belongs_to), calls s.Delete, then cascades to the captured parents. Share the visited map and recursion helper with UpdateCascade.
id: TSK-0037
kind: task
name: Cascade on Create and Delete paths
objective: Creating or removing an entity also bumps its parent's UpdatedAt so the parent reflects the child-set change
plan_phase: phase_3
plan_ref: task-done-affected-flags-6zl4
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: cascade-on-create-and-delete-paths-mbhq
task_status: completed
updated_at: "2026-04-14T07:31:57Z"
---
