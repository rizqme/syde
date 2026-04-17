---
id: CON-0059
kind: contract
name: Live Updates WebSocket
slug: live-updates-websocket-hev6
description: WS /api/<project>/ws — server-pushed change events for live SPA refresh.
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: websocket-server
      type: references
updated_at: "2026-04-16T10:51:15Z"
contract_kind: websocket
interaction_pattern: pub-sub
input: WS /api/<project>/ws
input_parameters:
    - path: project
      type: string
      description: path parameter
output: Server-pushed change events
output_parameters:
    - path: type
      type: string
      description: event type e.g. entity_changed, tree_changed
    - path: payload
      type: object
      description: type-dependent change details
---
