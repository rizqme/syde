---
id: REQ-0293
kind: requirement
name: Relationship Graph Invocation
slug: relationship-graph-invocation-zft1
relationships:
    - target: relationship-graph-erzs
      type: refines
    - target: graph-engine-xgjy
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:39Z"
statement: When the user runs syde graph, the syde CLI shall render an ASCII tree or DOT text representation of the entity relationship graph on stdout.
req_type: interface
priority: must
verification: integration test invoking syde graph with ascii and dot formats
source: manual
source_ref: contract:relationship-graph-erzs
requirement_status: active
rationale: Graph rendering is the primary way operators visualize how entities relate.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:39Z"
    graph-engine-xgjy:
        hash: 008188a7a397c93a8d847fa561e5274e77480780d36faacff440a814f6d605fe
        at: "2026-04-18T09:36:39Z"
---
