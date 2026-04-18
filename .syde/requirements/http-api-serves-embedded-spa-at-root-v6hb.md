---
id: REQ-0047
kind: requirement
name: HTTP API Serves Embedded SPA At Root
slug: http-api-serves-embedded-spa-at-root-v6hb
relationships:
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:59Z"
statement: When a client invokes GET /, the syded daemon shall respond with the embedded React SPA index document.
req_type: interface
priority: must
verification: integration test against / expecting text/html
source: manual
source_ref: component:http-api-afos
requirement_status: active
rationale: The daemon is a single binary that bundles the UI; the root must serve it.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:59Z"
---
