---
id: REQ-0129
kind: requirement
name: Daemon Launcher Probes Health Endpoint
slug: daemon-launcher-probes-health-endpoint-iyhi
relationships:
    - target: daemon-launcher-tzso
      type: refines
updated_at: "2026-04-18T09:37:37Z"
statement: When EnsureRunning is called, the daemon launcher shall probe the syded /health endpoint with a short timeout to detect a running instance.
req_type: functional
priority: must
verification: inspection of EnsureRunning in internal/daemon/daemon.go
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: The probe avoids spawning duplicate syded processes.
verified_against:
    daemon-launcher-tzso:
        hash: 6eff903308820d498bcecf7735f691f70ee70cb5fd4b94866e01c29ffa0a9645
        at: "2026-04-18T09:37:37Z"
---
