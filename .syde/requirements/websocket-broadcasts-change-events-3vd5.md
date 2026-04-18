---
id: REQ-0101
kind: requirement
name: WebSocket Broadcasts Change Events
slug: websocket-broadcasts-change-events-3vd5
relationships:
    - target: websocket-server-hdup
      type: refines
updated_at: "2026-04-18T09:37:05Z"
statement: The syded WebSocket server shall broadcast entity and tree change events to connected dashboard clients.
req_type: functional
priority: must
verification: end-to-end test publishing a change event and asserting the client receives it
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: Live updates are the entire purpose of the WebSocket server.
verified_against:
    websocket-server-hdup:
        hash: c2d6b7c8e035720d03d5754cdb196199ca2da6a30d1db538fa95a0ba7fe31fc5
        at: "2026-04-18T09:37:05Z"
---
