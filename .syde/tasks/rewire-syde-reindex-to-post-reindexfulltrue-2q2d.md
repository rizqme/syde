---
acceptance: syde reindex still reports counts and clean-rebuilds the on-disk index
completed_at: "2026-04-14T06:58:09Z"
details: 'internal/cli/reindex.go: client.Reindex with full:true; print the returned stats.'
id: TSK-0025
kind: task
name: Rewire syde reindex to POST /reindex?full=true
objective: syde reindex becomes a thin RPC triggering a full rebuild server-side
plan_phase: phase_4
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: rewire-syde-reindex-to-post-reindexfulltrue-2q2d
task_status: completed
updated_at: "2026-04-14T06:58:09Z"
---
