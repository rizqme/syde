---
id: REQ-0107
kind: requirement
name: CLI HTTP Client Does Not Access BadgerDB
slug: cli-http-client-does-not-access-badgerdb-rjl1
relationships:
    - target: cli-http-client-otp2
      type: refines
updated_at: "2026-04-18T09:37:19Z"
statement: The CLI HTTP client shall not access BadgerDB directly and shall route every read and write through syded HTTP endpoints.
req_type: constraint
priority: must
verification: code review of internal/client for storage imports
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Single-writer discipline keeps the index consistent.
verified_against:
    cli-http-client-otp2:
        hash: a871d2841c81ac5569c786e9c7a17276e38812f9aee8870bb11d15c8e3cb3d54
        at: "2026-04-18T09:37:19Z"
---
