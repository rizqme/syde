---
acceptance: 'Unit-test-style: updating a component bumps the system it belongs to; updating that system is a no-op (no parent); a fabricated cycle A->B->A terminates after two Updates (one each)'
affected_entities:
    - storage-engine
affected_files:
    - internal/storage/store.go
completed_at: "2026-04-14T07:31:57Z"
created_at: "2026-04-14T07:26:27Z"
details: Implement s.cascadeFromParent(base, visited) and s.UpdateCascade(e, body) in internal/storage/store.go. visited is map[string]bool keyed by entity ID. The child's own ID is added to visited before recursing so self-loops terminate. Each parent traversal calls s.Get to resolve the slug (handles ID, full slug, bare slug, parent/child path), skips on Get error (broken ref), marks visited, then s.Update + recurse.
id: TSK-0036
kind: task
name: Add cascadeFromParent helper + UpdateCascade on Store
objective: Store can update an entity and recursively bump its belongs_to ancestors
plan_phase: phase_3
plan_ref: task-done-affected-flags-6zl4
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: add-cascadefromparent-helper-updatecascade-on-store-6sxq
task_status: completed
updated_at: "2026-04-14T07:31:57Z"
---
