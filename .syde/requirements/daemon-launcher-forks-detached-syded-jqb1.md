---
id: REQ-0133
kind: requirement
name: Daemon Launcher Forks Detached syded
slug: daemon-launcher-forks-detached-syded-jqb1
relationships:
    - target: daemon-launcher-tzso
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:58Z"
statement: When no running syded is detected, the daemon launcher shall fork a detached syded process with daemon flags and idle timeout so it survives the CLI exit.
req_type: functional
priority: must
verification: integration test that CLI exit leaves syded running
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Auto-started syded must outlive the short-lived CLI invocation.
---
