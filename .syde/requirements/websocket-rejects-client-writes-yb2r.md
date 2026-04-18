---
id: REQ-0104
kind: requirement
name: WebSocket Rejects Client Writes
slug: websocket-rejects-client-writes-yb2r
relationships:
    - target: websocket-server-hdup
      type: refines
updated_at: "2026-04-18T09:37:45Z"
statement: If a client sends a write or mutation frame over the WebSocket, then the syded WebSocket server shall ignore the frame and shall not apply any mutation.
req_type: constraint
priority: must
verification: end-to-end test sending a write frame and asserting no state change
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: The dashboard is read-only; accepting writes over the socket would bypass that invariant.
verified_against:
    websocket-server-hdup:
        hash: c2d6b7c8e035720d03d5754cdb196199ca2da6a30d1db538fa95a0ba7fe31fc5
        at: "2026-04-18T09:37:45Z"
---
