---
id: REQ-0110
kind: requirement
name: CLI HTTP Client Does Not Cache Results
slug: cli-http-client-does-not-cache-results-d6dc
relationships:
    - target: cli-http-client-otp2
      type: refines
updated_at: "2026-04-18T09:37:55Z"
statement: The CLI HTTP client shall not cache response bodies and shall issue a fresh HTTP request for each call.
req_type: constraint
priority: must
verification: code review of client.go for caching layers
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Cache-free behavior ensures CLI output always reflects the current store state.
verified_against:
    cli-http-client-otp2:
        hash: a871d2841c81ac5569c786e9c7a17276e38812f9aee8870bb11d15c8e3cb3d54
        at: "2026-04-18T09:37:55Z"
---
