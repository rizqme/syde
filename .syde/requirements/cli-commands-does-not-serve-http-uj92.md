---
id: REQ-0091
kind: requirement
name: CLI Commands Does Not Serve HTTP
slug: cli-commands-does-not-serve-http-uj92
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:01Z"
statement: The syde CLI shall not serve HTTP endpoints or render HTML and shall delegate dashboard concerns to syded.
req_type: constraint
priority: must
verification: code review of cmd/syde for HTTP server imports
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: The syde binary is a client; syded owns the server surface.
---
