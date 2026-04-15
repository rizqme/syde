---
id: REQ-0289
kind: requirement
name: Query Entity Invocation
slug: query-entity-invocation-wt5d
relationships:
    - target: query-entity-ci7d
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde query <slug>, the syde CLI shall print a formatted query result for the named entity on stdout.
req_type: interface
priority: must
verification: integration test invoking syde query on an existing slug
source: manual
source_ref: contract:query-entity-ci7d
requirement_status: active
rationale: syde query is the canonical rich accessor used for deep dives and impact analysis.
---
