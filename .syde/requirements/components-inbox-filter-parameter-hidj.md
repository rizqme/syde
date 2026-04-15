---
id: REQ-0217
kind: requirement
name: Components Inbox Filter Parameter
slug: components-inbox-filter-parameter-hidj
relationships:
    - target: components-inbox-screen-c5jh
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:27Z"
statement: Where a filter query parameter is provided on the /component route, the dashboard shall restrict the components list to entities matching the filter DSL query.
req_type: interface
priority: should
verification: manual inspection of /component?filter=... in the dashboard
source: manual
source_ref: contract:components-inbox-screen-c5jh
requirement_status: active
rationale: Filter support keeps large component inventories navigable.
---
