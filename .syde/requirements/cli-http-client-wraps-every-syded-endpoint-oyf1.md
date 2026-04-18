---
id: REQ-0095
kind: requirement
name: CLI HTTP Client Wraps Every syded Endpoint
slug: cli-http-client-wraps-every-syded-endpoint-oyf1
relationships:
    - target: cli-http-client-otp2
      type: refines
updated_at: "2026-04-18T09:38:01Z"
statement: The CLI HTTP client shall wrap every syded read and write endpoint as a typed Go method callable from CLI commands.
req_type: functional
priority: must
verification: inspection of internal/client/client.go method coverage
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: A complete wrapper gives the CLI a zero-BadgerDB read and write path.
verified_against:
    cli-http-client-otp2:
        hash: a871d2841c81ac5569c786e9c7a17276e38812f9aee8870bb11d15c8e3cb3d54
        at: "2026-04-18T09:38:01Z"
---
