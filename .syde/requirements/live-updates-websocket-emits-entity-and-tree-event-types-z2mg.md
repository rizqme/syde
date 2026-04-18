---
id: REQ-0254
kind: requirement
name: Live Updates WebSocket Emits Entity And Tree Event Types
slug: live-updates-websocket-emits-entity-and-tree-event-types-z2mg
relationships:
    - target: live-updates-websocket-hev6
      type: refines
    - target: http-api-afos
      type: refines
    - target: websocket-server-hdup
      type: refines
updated_at: "2026-04-18T09:36:43Z"
statement: When the syded daemon emits events over /api/<project>/ws, the syded daemon shall tag each event with a type field identifying the change category such as entity_changed or tree_changed.
req_type: interface
priority: must
verification: end-to-end test publishing entity_changed and tree_changed events over /api/<project>/ws
source: manual
source_ref: contract:live-updates-websocket-hev6
requirement_status: active
rationale: Clients demultiplex change handling by event type.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:36:43Z"
    websocket-server-hdup:
        hash: c2d6b7c8e035720d03d5754cdb196199ca2da6a30d1db538fa95a0ba7fe31fc5
        at: "2026-04-18T09:36:43Z"
---
