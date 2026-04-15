---
id: REQ-0212
kind: requirement
name: Systems Inbox Filter Parameter
slug: systems-inbox-filter-parameter-n2kf
relationships:
    - target: systems-inbox-screen-qgp4
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:22Z"
statement: Where a filter query parameter is provided on the /system route, the dashboard shall restrict the systems list to entities matching the filter DSL query.
req_type: interface
priority: should
verification: manual inspection of /system?filter=... in the dashboard
source: manual
source_ref: contract:systems-inbox-screen-qgp4
requirement_status: active
rationale: Filter support lets users scope large inboxes to a workable subset.
---
