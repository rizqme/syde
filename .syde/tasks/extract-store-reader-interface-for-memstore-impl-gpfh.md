---
acceptance: go build passes; existing dashboard tests still pass
affected_files:
    - internal/storage/store.go
completed_at: "2026-04-14T06:23:18Z"
created_at: "2026-04-14T06:03:22Z"
details: Define ReadStore interface with Get/List/Query/GetInbound/etc. Make badger-backed Store satisfy it. Update dashboard handlers to take ReadStore.
id: TSK-0005
kind: task
name: Extract Store reader interface for memstore impl
objective: Dashboard handlers depend on an interface, not concrete badger-backed struct
plan_phase: phase_4
plan_ref: cli-health-daemon-coexistence-p25w
relationships:
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - target: syde-5tdt
      type: belongs_to
slug: extract-store-reader-interface-for-memstore-impl-gpfh
task_status: completed
updated_at: "2026-04-14T06:23:18Z"
---
