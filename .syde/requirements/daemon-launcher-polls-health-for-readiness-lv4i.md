---
id: REQ-0137
kind: requirement
name: Daemon Launcher Polls Health For Readiness
slug: daemon-launcher-polls-health-for-readiness-lv4i
relationships:
    - target: daemon-launcher-tzso
      type: refines
updated_at: "2026-04-18T09:36:45Z"
statement: When the daemon launcher is waiting for a freshly spawned syded, the daemon launcher shall poll the health endpoint every 50 milliseconds for up to 3 seconds.
req_type: performance
priority: must
verification: inspection of readiness polling in daemon.go
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Bounded readiness polling gives fast feedback while covering typical startup latency.
verified_against:
    daemon-launcher-tzso:
        hash: 6eff903308820d498bcecf7735f691f70ee70cb5fd4b94866e01c29ffa0a9645
        at: "2026-04-18T09:36:45Z"
---
