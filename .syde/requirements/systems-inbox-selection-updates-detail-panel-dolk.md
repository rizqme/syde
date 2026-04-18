---
id: REQ-0210
kind: requirement
name: Systems Inbox Selection Updates Detail Panel
slug: systems-inbox-selection-updates-detail-panel-dolk
relationships:
    - target: systems-inbox-screen-qgp4
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:38:11Z"
statement: When the user clicks a system in the systems inbox list, the dashboard shall display that system's details in the detail panel.
req_type: interface
priority: must
verification: manual inspection of /system in the dashboard
source: manual
source_ref: contract:systems-inbox-screen-qgp4
requirement_status: active
rationale: Selection-to-detail is the core interaction of the two-column inbox pattern.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:38:11Z"
---
