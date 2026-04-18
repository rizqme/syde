---
id: REQ-0217
kind: requirement
name: Components Inbox Filter Parameter
slug: components-inbox-filter-parameter-hidj
relationships:
    - target: components-inbox-screen-c5jh
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:36:59Z"
statement: Where a filter query parameter is provided on the /component route, the dashboard shall restrict the components list to entities matching the filter DSL query.
req_type: interface
priority: should
verification: manual inspection of /component?filter=... in the dashboard
source: manual
source_ref: contract:components-inbox-screen-c5jh
requirement_status: active
rationale: Filter support keeps large component inventories navigable.
audited_overlaps:
    - slug: systems-inbox-filter-parameter-n2kf
      distinction: 'Different route and entity kind: this filters the /component inbox for components, while the other filters the /system inbox for systems.'
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:36:59Z"
---
