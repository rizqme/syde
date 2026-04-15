---
contract_kind: websocket
description: WS /api/<project>/ws — server-pushed change events for live SPA refresh.
id: CON-0059
input: WS /api/<project>/ws
input_parameters:
    - description: path parameter
      path: project
      type: string
interaction_pattern: pub-sub
kind: contract
name: Live Updates WebSocket
output: Server-pushed change events
output_parameters:
    - description: event type e.g. entity_changed, tree_changed
      path: type
      type: string
    - description: type-dependent change details
      path: payload
      type: object
relationships:
    - target: syded-dashboard
      type: belongs_to
    - target: websocket-server
      type: references
    - label: requirement
      target: existing-syde-model-baseline-hcvj
      type: references
    - label: flow
      target: design-model-operations-coverage-wsrh
      type: involves
slug: live-updates-websocket-hev6
updated_at: "2026-04-14T03:27:06Z"
---
