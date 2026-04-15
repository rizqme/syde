---
id: REQ-0252
kind: requirement
name: Live Updates WebSocket Pushes Change Events
slug: live-updates-websocket-pushes-change-events-otz2
relationships:
    - target: live-updates-websocket-hev6
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T11:00:06Z"
statement: When a client opens WS /api/<project>/ws, the syded daemon shall push server-initiated change event messages containing a type string and a payload object for the duration of the connection.
req_type: interface
priority: must
verification: end-to-end test publishing a change event over /api/<project>/ws
source: manual
source_ref: contract:live-updates-websocket-hev6
requirement_status: active
rationale: The SPA live-refreshes in response to server-pushed changes.
---
