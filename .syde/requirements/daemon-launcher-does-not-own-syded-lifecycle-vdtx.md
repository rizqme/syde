---
id: REQ-0135
kind: requirement
name: Daemon Launcher Does Not Own syded Lifecycle
slug: daemon-launcher-does-not-own-syded-lifecycle-vdtx
relationships:
    - target: daemon-launcher-tzso
      type: refines
updated_at: "2026-04-18T09:37:35Z"
statement: The daemon launcher shall not own the full syded lifecycle beyond the initial spawn and readiness check.
req_type: constraint
priority: must
verification: code review of internal/daemon for restart logic
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Restart responsibility belongs to an external supervisor, not the CLI-side launcher.
verified_against:
    daemon-launcher-tzso:
        hash: 6eff903308820d498bcecf7735f691f70ee70cb5fd4b94866e01c29ffa0a9645
        at: "2026-04-18T09:37:35Z"
---
