---
id: REQ-0106
kind: requirement
name: CLI HTTP Client Sends Write Payloads As YAML Frontmatter
slug: cli-http-client-sends-write-payloads-as-yaml-frontmatter-94wy
relationships:
    - target: cli-http-client-otp2
      type: refines
updated_at: "2026-04-18T09:36:50Z"
statement: When the CLI HTTP client invokes a create or update entity endpoint, the client shall send the entity body as a YAML frontmatter payload.
req_type: interface
priority: must
verification: inspection of CreateEntity/UpdateEntity in client.go
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: The syded server expects the markdown-on-the-wire format for round-trippable persistence.
verified_against:
    cli-http-client-otp2:
        hash: a871d2841c81ac5569c786e9c7a17276e38812f9aee8870bb11d15c8e3cb3d54
        at: "2026-04-18T09:36:50Z"
---
