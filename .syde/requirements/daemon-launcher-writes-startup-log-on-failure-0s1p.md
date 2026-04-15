---
id: REQ-0134
kind: requirement
name: Daemon Launcher Writes Startup Log On Failure
slug: daemon-launcher-writes-startup-log-on-failure-0s1p
relationships:
    - target: daemon-launcher-tzso
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:03Z"
statement: If readiness polling fails, then the daemon launcher shall report the tail of the syded startup log from the user home directory to the caller.
req_type: functional
priority: must
verification: integration test that failed readiness surfaces log tail
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Users need diagnostic output when auto-start cannot reach a healthy daemon.
---
