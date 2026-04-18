---
id: REQ-0091
kind: requirement
name: CLI Commands Does Not Serve HTTP
slug: cli-commands-does-not-serve-http-uj92
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:33Z"
statement: The syde CLI shall not serve HTTP endpoints or render HTML and shall delegate dashboard concerns to syded.
req_type: constraint
priority: must
verification: code review of cmd/syde for HTTP server imports
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: The syde binary is a client; syded owns the server surface.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:33Z"
---
