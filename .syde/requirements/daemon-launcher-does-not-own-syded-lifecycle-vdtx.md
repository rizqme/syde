---
id: REQ-0135
kind: requirement
name: Daemon Launcher Does Not Own syded Lifecycle
slug: daemon-launcher-does-not-own-syded-lifecycle-vdtx
relationships:
    - target: daemon-launcher-tzso
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:05Z"
statement: The daemon launcher shall not own the full syded lifecycle beyond the initial spawn and readiness check.
req_type: constraint
priority: must
verification: code review of internal/daemon for restart logic
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Restart responsibility belongs to an external supervisor, not the CLI-side launcher.
---
