---
acceptance: Whole flow runs with zero manual syded starts and no lock errors
completed_at: "2026-04-14T07:00:44Z"
details: 'Run in /tmp/syde-smoke: syde init; syde add component Foo; syde query foo; syde sync check --strict; kill syded; syde status (should auto-relaunch). Document the sequence in the plan notes.'
id: TSK-0029
kind: task
name: 'End-to-end smoke test: CLI + syded coexist'
objective: Fresh clone -> init -> add -> query -> validate cycle works with syded auto-launched
plan_phase: phase_5
plan_ref: cli-syded-client-refactor-ozvj
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: end-to-end-smoke-test-cli-syded-coexist-aabd
task_status: completed
updated_at: "2026-04-14T07:00:44Z"
---
