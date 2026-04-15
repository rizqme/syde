---
id: REQ-0015
kind: requirement
name: syded shall watch markdown for live updates
slug: syded-shall-watch-markdown-for-live-updates-im0r
relationships:
    - target: syded-dashboard-e82c
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:49:02Z"
statement: When any file under .syde changes, the syded daemon shall push a change event to connected WebSocket clients within one second.
req_type: performance
priority: should
verification: manual test editing a markdown file while the dashboard is open
source: manual
source_ref: system:syded-dashboard-e82c:scope
requirement_status: active
rationale: Live updates make the dashboard usable as a side panel during CLI work.
---
