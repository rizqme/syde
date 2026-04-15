---
id: REQ-0102
kind: requirement
name: WebSocket Per Project Rooms
slug: websocket-per-project-rooms-rcdw
relationships:
    - target: websocket-server-hdup
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:16Z"
statement: The syded WebSocket server shall isolate broadcasts to per-project rooms so events from one project are not delivered to another.
req_type: functional
priority: must
verification: end-to-end test connecting two clients to different projects and asserting no cross-delivery
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: Cross-project leakage would surprise users running multiple projects on one daemon.
---
