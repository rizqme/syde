---
id: TSK-0261
kind: task
name: Build + reinstall + restart daemon
slug: build-reinstall-restart-daemon-pqbd
relationships:
    - target: syde-5tdt
      type: belongs_to
    - target: design-model-shall-contain-exactly-two-systems-named-syde-and-syded-f2q8
      type: implements
updated_at: "2026-04-18T09:44:54Z"
task_status: completed
priority: high
objective: Rebuilt syde/syded binaries with new audit rules + frontend; skill reinstalled; daemon restarted.
acceptance: make install succeeds; syded restarts cleanly; /health returns 200.
plan_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
plan_phase: phase_5
created_at: "2026-04-18T09:09:09Z"
completed_at: "2026-04-18T09:36:31Z"
---
