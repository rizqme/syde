---
id: REQ-0104
kind: requirement
name: WebSocket Rejects Client Writes
slug: websocket-rejects-client-writes-yb2r
relationships:
    - target: websocket-server-hdup
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:16Z"
statement: If a client sends a write or mutation frame over the WebSocket, then the syded WebSocket server shall ignore the frame and shall not apply any mutation.
req_type: constraint
priority: must
verification: end-to-end test sending a write frame and asserting no state change
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: The dashboard is read-only; accepting writes over the socket would bypass that invariant.
---
