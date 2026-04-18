---
id: REQ-0212
kind: requirement
name: Systems Inbox Filter Parameter
slug: systems-inbox-filter-parameter-n2kf
relationships:
    - target: systems-inbox-screen-qgp4
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:01Z"
statement: Where a filter query parameter is provided on the /system route, the dashboard shall restrict the systems list to entities matching the filter DSL query.
req_type: interface
priority: should
verification: manual inspection of /system?filter=... in the dashboard
source: manual
source_ref: contract:systems-inbox-screen-qgp4
requirement_status: active
rationale: Filter support lets users scope large inboxes to a workable subset.
audited_overlaps:
    - slug: components-inbox-filter-parameter-hidj
      distinction: Filter parameter governs the /system route systems list; the paired requirement governs the /component route components list.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:01Z"
---
