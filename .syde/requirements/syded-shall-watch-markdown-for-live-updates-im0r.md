---
id: REQ-0015
kind: requirement
name: syded shall watch markdown for live updates
slug: syded-shall-watch-markdown-for-live-updates-im0r
relationships:
    - target: http-api-afos
      type: refines
    - target: websocket-server-hdup
      type: refines
updated_at: "2026-04-18T09:36:53Z"
statement: When any file under .syde changes, the syded daemon shall push a change event to connected WebSocket clients within one second.
req_type: performance
priority: should
verification: manual test editing a markdown file while the dashboard is open
source: manual
source_ref: system:syded-dashboard-e82c:scope
requirement_status: active
rationale: Live updates make the dashboard usable as a side panel during CLI work.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:36:53Z"
    websocket-server-hdup:
        hash: c2d6b7c8e035720d03d5754cdb196199ca2da6a30d1db538fa95a0ba7fe31fc5
        at: "2026-04-18T09:36:53Z"
---
