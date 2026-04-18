---
id: REQ-0114
kind: requirement
name: CLI HTTP Client Supports Reindex RPC
slug: cli-http-client-supports-reindex-rpc-fwpt
relationships:
    - target: cli-http-client-otp2
      type: refines
updated_at: "2026-04-18T09:37:45Z"
statement: When the CLI invokes reindex, the CLI HTTP client shall issue a POST to the reindex endpoint and return the resulting counts.
req_type: interface
priority: must
verification: integration test against POST /reindex
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Reindex needs to be exposed through the same HTTP path as other writes.
verified_against:
    cli-http-client-otp2:
        hash: a871d2841c81ac5569c786e9c7a17276e38812f9aee8870bb11d15c8e3cb3d54
        at: "2026-04-18T09:37:45Z"
---
