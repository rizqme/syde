---
id: REQ-0131
kind: requirement
name: Daemon Launcher Locates syded Binary
slug: daemon-launcher-locates-syded-binary-tmv8
relationships:
    - target: daemon-launcher-tzso
      type: refines
updated_at: "2026-04-18T09:37:43Z"
statement: The daemon launcher shall locate the syded binary via the SYDED_BIN environment variable, a sibling of the current executable, or the system PATH.
req_type: functional
priority: must
verification: inspection of binary lookup in daemon.go
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Flexible binary discovery supports dev and production installs.
verified_against:
    daemon-launcher-tzso:
        hash: 6eff903308820d498bcecf7735f691f70ee70cb5fd4b94866e01c29ffa0a9645
        at: "2026-04-18T09:37:43Z"
---
