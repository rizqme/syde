---
id: REQ-0137
kind: requirement
name: Daemon Launcher Polls Health For Readiness
slug: daemon-launcher-polls-health-for-readiness-lv4i
relationships:
    - target: daemon-launcher-tzso
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:14Z"
statement: When the daemon launcher is waiting for a freshly spawned syded, the daemon launcher shall poll the health endpoint every 50 milliseconds for up to 3 seconds.
req_type: performance
priority: must
verification: inspection of readiness polling in daemon.go
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Bounded readiness polling gives fast feedback while covering typical startup latency.
---
