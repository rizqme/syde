---
id: REQ-0293
kind: requirement
name: Relationship Graph Invocation
slug: relationship-graph-invocation-zft1
relationships:
    - target: relationship-graph-erzs
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:01:33Z"
statement: When the user runs syde graph, the syde CLI shall render an ASCII tree or DOT text representation of the entity relationship graph on stdout.
req_type: interface
priority: must
verification: integration test invoking syde graph with ascii and dot formats
source: manual
source_ref: contract:relationship-graph-erzs
requirement_status: active
rationale: Graph rendering is the primary way operators visualize how entities relate.
---
