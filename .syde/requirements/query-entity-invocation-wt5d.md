---
id: REQ-0289
kind: requirement
name: Query Entity Invocation
slug: query-entity-invocation-wt5d
relationships:
    - target: query-entity-ci7d
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:34Z"
statement: When the user runs syde query <slug>, the syde CLI shall print a formatted query result for the named entity on stdout.
req_type: interface
priority: must
verification: integration test invoking syde query on an existing slug
source: manual
source_ref: contract:query-entity-ci7d
requirement_status: active
rationale: syde query is the canonical rich accessor used for deep dives and impact analysis.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:34Z"
---
