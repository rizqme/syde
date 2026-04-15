---
id: REQ-0210
kind: requirement
name: Systems Inbox Selection Updates Detail Panel
slug: systems-inbox-selection-updates-detail-panel-dolk
relationships:
    - target: systems-inbox-screen-qgp4
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:19Z"
statement: When the user clicks a system in the systems inbox list, the dashboard shall display that system's details in the detail panel.
req_type: interface
priority: must
verification: manual inspection of /system in the dashboard
source: manual
source_ref: contract:systems-inbox-screen-qgp4
requirement_status: active
rationale: Selection-to-detail is the core interaction of the two-column inbox pattern.
---
