---
id: REQ-0218
kind: requirement
name: List Entities Accepts Tag Filter
slug: list-entities-accepts-tag-filter-gdem
tags:
    - flow-coverage-reviewed
relationships:
    - target: list-entities-http-t17z
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-17T10:55:11Z"
statement: When GET /api/<project>/entities is invoked, the syded daemon shall accept an optional tag query parameter as a string and restrict results to entities carrying that tag.
req_type: interface
priority: must
verification: integration test against /api/<project>/entities?tag=core
source: manual
source_ref: contract:list-entities-http-t17z
requirement_status: active
rationale: Dashboard pages filter entity lists by tag.
audited_overlaps:
    - slug: list-entities-accepts-kind-filter-rjf6
      distinction: Filters by tag query parameter; target filters by entity kind query parameter on the same endpoint.
---
