---
id: REQ-0072
kind: requirement
name: CLI Commands Support Query Graph And Search
slug: cli-commands-support-query-graph-and-search-jvcq
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:46Z"
statement: The syde CLI shall expose query, search, and graph subcommands for navigating the entity model.
req_type: functional
priority: must
verification: integration test invoking syde query, search, graph
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: Read paths drive context discovery for agents and humans.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:46Z"
---
