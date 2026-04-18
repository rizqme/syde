---
id: TSK-0262
kind: task
name: Bulk re-verify all active requirements after audit rule change
slug: bulk-re-verify-all-active-requirements-after-audit-rule-change-rdba
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: component-shall-be-allowed-to-belong-to-multiple-systems-qd6u
      type: implements
updated_at: "2026-04-18T09:44:54Z"
task_status: completed
priority: high
objective: Every active requirement carries a fresh verified_against snapshot reflecting post-Phase-1 component file hashes.
acceptance: syde sync check reports zero requirement_stale findings.
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_5
created_at: "2026-04-18T09:09:10Z"
completed_at: "2026-04-18T09:38:17Z"
---
