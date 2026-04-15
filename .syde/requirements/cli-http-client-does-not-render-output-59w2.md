---
id: REQ-0112
kind: requirement
name: CLI HTTP Client Does Not Render Output
slug: cli-http-client-does-not-render-output-59w2
relationships:
    - target: cli-http-client-otp2
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:27Z"
statement: The CLI HTTP client shall not render output for terminals and shall return decoded Go structs to its callers.
req_type: constraint
priority: must
verification: inspection of client.go return types
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Formatting is the CLI command layer's responsibility.
---
