---
id: REQ-0254
kind: requirement
name: Live Updates WebSocket Emits Entity And Tree Event Types
slug: live-updates-websocket-emits-entity-and-tree-event-types-z2mg
relationships:
    - target: live-updates-websocket-hev6
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T11:00:13Z"
statement: When the syded daemon emits events over /api/<project>/ws, the syded daemon shall tag each event with a type field identifying the change category such as entity_changed or tree_changed.
req_type: interface
priority: must
verification: end-to-end test publishing entity_changed and tree_changed events over /api/<project>/ws
source: manual
source_ref: contract:live-updates-websocket-hev6
requirement_status: active
rationale: Clients demultiplex change handling by event type.
---
