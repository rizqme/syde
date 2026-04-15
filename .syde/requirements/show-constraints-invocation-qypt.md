---
id: REQ-0301
kind: requirement
name: Show Constraints Invocation
slug: show-constraints-invocation-qypt
relationships:
    - target: show-constraints-jnig
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde constraints, the syde CLI shall print all active architecture decisions on stdout.
req_type: interface
priority: must
verification: integration test invoking syde constraints
source: manual
source_ref: contract:show-constraints-jnig
requirement_status: active
rationale: Listing active decisions lets operators review the full constraint surface at once.
---
