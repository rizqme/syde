---
id: TSK-0258
kind: task
name: Bulk-apply component and contract reparenting and reindex
slug: bulk-apply-component-and-contract-reparenting-and-reindex-gqdc
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: component-shall-be-allowed-to-belong-to-multiple-systems-qd6u
      type: implements
updated_at: "2026-04-18T09:44:53Z"
task_status: completed
priority: high
objective: All 12 components and 57 contracts have belongs_to edges pointing to syde-5tdt and/or syded-dashboard-e82c per the worksheet; none still references syde-cli-2478.
details: Python script reads each target YAML, drops relationships with type=belongs_to and target∈{syde-cli, syde-cli-2478}, appends the proposed belongs_to edges from the worksheet, writes back. Then syde reindex.
acceptance: syde reindex succeeds; spot-check 5 random entities show no belongs_to:syde-cli-2478 and the expected new targets.
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_3
created_at: "2026-04-18T09:09:09Z"
completed_at: "2026-04-18T09:30:47Z"
---
