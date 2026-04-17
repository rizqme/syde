---
id: TSK-0021
kind: task
name: End-to-end plan smoke test
slug: end-to-end-plan-smoke-test-c87y
relationships:
    - target: revamp-planning-to-structured-design-and-diff
      type: belongs_to
    - target: plans-shall-pass-syde-plan-check-before-approval-0jkc
      type: references
updated_at: "2026-04-16T01:33:06Z"
task_status: completed
priority: high
objective: go test + make install + full plan lifecycle smoke test all pass.
details: 'go test ./..., make install, syde install-skill --all. Then in /tmp/syde-plan-smoke: syde init, create two components, create a plan with --design and three changes (one Extended FieldChanges on a component, one New contract with contract_kind screen + wireframe, one Deleted requirement), approve, add a phase and a task, complete the task, run syde plan complete and expect failure on deliberately unchanged extended field, fix the field with syde update, retry and expect success.'
acceptance: Smoke script runs end-to-end with the validator catching the deliberate mismatch and accepting the corrected state.
affected_entities:
    - audit-engine-4ktg
    - cli-commands-hpjb
plan_ref: revamp-planning-to-structured-design-and-diff-m8p5
plan_phase: phase_6
created_at: "2026-04-15T11:42:14Z"
completed_at: "2026-04-15T12:42:19Z"
---
