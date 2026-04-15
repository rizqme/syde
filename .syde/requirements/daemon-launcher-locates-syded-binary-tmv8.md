---
id: REQ-0131
kind: requirement
name: Daemon Launcher Locates syded Binary
slug: daemon-launcher-locates-syded-binary-tmv8
relationships:
    - target: daemon-launcher-tzso
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:55Z"
statement: The daemon launcher shall locate the syded binary via the SYDED_BIN environment variable, a sibling of the current executable, or the system PATH.
req_type: functional
priority: must
verification: inspection of binary lookup in daemon.go
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Flexible binary discovery supports dev and production installs.
---
