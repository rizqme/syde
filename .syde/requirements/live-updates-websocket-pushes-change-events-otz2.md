---
id: REQ-0252
kind: requirement
name: Live Updates WebSocket Pushes Change Events
slug: live-updates-websocket-pushes-change-events-otz2
relationships:
    - target: live-updates-websocket-hev6
      type: refines
    - target: http-api-afos
      type: refines
    - target: websocket-server-hdup
      type: refines
updated_at: "2026-04-18T09:37:11Z"
statement: When a client opens WS /api/<project>/ws, the syded daemon shall push server-initiated change event messages containing a type string and a payload object for the duration of the connection.
req_type: interface
priority: must
verification: end-to-end test publishing a change event over /api/<project>/ws
source: manual
source_ref: contract:live-updates-websocket-hev6
requirement_status: active
rationale: The SPA live-refreshes in response to server-pushed changes.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:11Z"
    websocket-server-hdup:
        hash: c2d6b7c8e035720d03d5754cdb196199ca2da6a30d1db538fa95a0ba7fe31fc5
        at: "2026-04-18T09:37:11Z"
---
