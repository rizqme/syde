---
id: REQ-0095
kind: requirement
name: CLI HTTP Client Wraps Every syded Endpoint
slug: cli-http-client-wraps-every-syded-endpoint-oyf1
relationships:
    - target: cli-http-client-otp2
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:10Z"
statement: The CLI HTTP client shall wrap every syded read and write endpoint as a typed Go method callable from CLI commands.
req_type: functional
priority: must
verification: inspection of internal/client/client.go method coverage
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: A complete wrapper gives the CLI a zero-BadgerDB read and write path.
---
