---
id: REQ-0114
kind: requirement
name: CLI HTTP Client Supports Reindex RPC
slug: cli-http-client-supports-reindex-rpc-fwpt
relationships:
    - target: cli-http-client-otp2
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:30Z"
statement: When the CLI invokes reindex, the CLI HTTP client shall issue a POST to the reindex endpoint and return the resulting counts.
req_type: interface
priority: must
verification: integration test against POST /reindex
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Reindex needs to be exposed through the same HTTP path as other writes.
---
