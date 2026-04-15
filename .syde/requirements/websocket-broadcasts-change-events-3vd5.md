---
id: REQ-0101
kind: requirement
name: WebSocket Broadcasts Change Events
slug: websocket-broadcasts-change-events-3vd5
relationships:
    - target: websocket-server-hdup
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:16Z"
statement: The syded WebSocket server shall broadcast entity and tree change events to connected dashboard clients.
req_type: functional
priority: must
verification: end-to-end test publishing a change event and asserting the client receives it
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: Live updates are the entire purpose of the WebSocket server.
---
