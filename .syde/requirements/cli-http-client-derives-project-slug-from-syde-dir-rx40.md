---
id: REQ-0097
kind: requirement
name: CLI HTTP Client Derives Project Slug From syde Dir
slug: cli-http-client-derives-project-slug-from-syde-dir-rx40
relationships:
    - target: cli-http-client-otp2
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:13Z"
statement: When the CLI HTTP client is constructed against a .syde directory, the client shall derive the project slug so requests route to the correct syded project.
req_type: functional
priority: must
verification: inspection of project slug derivation in client.go
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Multiple projects share a single syded; routing must be deterministic per directory.
---
