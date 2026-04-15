---
id: REQ-0107
kind: requirement
name: CLI HTTP Client Does Not Access BadgerDB
slug: cli-http-client-does-not-access-badgerdb-rjl1
relationships:
    - target: cli-http-client-otp2
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:20Z"
statement: The CLI HTTP client shall not access BadgerDB directly and shall route every read and write through syded HTTP endpoints.
req_type: constraint
priority: must
verification: code review of internal/client for storage imports
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Single-writer discipline keeps the index consistent.
---
