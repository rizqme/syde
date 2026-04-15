---
id: REQ-0047
kind: requirement
name: HTTP API Serves Embedded SPA At Root
slug: http-api-serves-embedded-spa-at-root-v6hb
relationships:
    - target: http-api-afos
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:22Z"
statement: When a client invokes GET /, the syded daemon shall respond with the embedded React SPA index document.
req_type: interface
priority: must
verification: integration test against / expecting text/html
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The daemon is a single binary that bundles the UI; the root must serve it.
---
