---
id: REQ-0134
kind: requirement
name: Daemon Launcher Writes Startup Log On Failure
slug: daemon-launcher-writes-startup-log-on-failure-0s1p
relationships:
    - target: daemon-launcher-tzso
      type: refines
updated_at: "2026-04-18T09:37:39Z"
statement: If readiness polling fails, then the daemon launcher shall report the tail of the syded startup log from the user home directory to the caller.
req_type: functional
priority: must
verification: integration test that failed readiness surfaces log tail
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Users need diagnostic output when auto-start cannot reach a healthy daemon.
verified_against:
    daemon-launcher-tzso:
        hash: 6eff903308820d498bcecf7735f691f70ee70cb5fd4b94866e01c29ffa0a9645
        at: "2026-04-18T09:37:39Z"
---
