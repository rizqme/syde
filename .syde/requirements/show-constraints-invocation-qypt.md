---
id: REQ-0301
kind: requirement
name: Show Constraints Invocation
slug: show-constraints-invocation-qypt
relationships:
    - target: show-constraints-jnig
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:30Z"
statement: When the user runs syde constraints, the syde CLI shall print all active architecture decisions on stdout.
req_type: interface
priority: must
verification: integration test invoking syde constraints
source: manual
source_ref: contract:show-constraints-jnig
requirement_status: active
rationale: Listing active decisions lets operators review the full constraint surface at once.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:30Z"
---
