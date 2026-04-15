---
id: REQ-0072
kind: requirement
name: CLI Commands Support Query Graph And Search
slug: cli-commands-support-query-graph-and-search-jvcq
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:46Z"
statement: The syde CLI shall expose query, search, and graph subcommands for navigating the entity model.
req_type: functional
priority: must
verification: integration test invoking syde query, search, graph
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: Read paths drive context discovery for agents and humans.
---
