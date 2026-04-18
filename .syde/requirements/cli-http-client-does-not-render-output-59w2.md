---
id: REQ-0112
kind: requirement
name: CLI HTTP Client Does Not Render Output
slug: cli-http-client-does-not-render-output-59w2
relationships:
    - target: cli-http-client-otp2
      type: refines
updated_at: "2026-04-18T09:37:33Z"
statement: The CLI HTTP client shall not render output for terminals and shall return decoded Go structs to its callers.
req_type: constraint
priority: must
verification: inspection of client.go return types
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Formatting is the CLI command layer's responsibility.
verified_against:
    cli-http-client-otp2:
        hash: a871d2841c81ac5569c786e9c7a17276e38812f9aee8870bb11d15c8e3cb3d54
        at: "2026-04-18T09:37:33Z"
---
