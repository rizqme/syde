---
id: REQ-0105
kind: requirement
name: WebSocket Not A Bulk Fetch Channel
slug: websocket-not-a-bulk-fetch-channel-g5tn
relationships:
    - target: websocket-server-hdup
      type: refines
updated_at: "2026-04-18T09:37:57Z"
statement: The syded WebSocket server shall not serve bulk entity fetches and shall defer bulk data retrieval to the HTTP API.
req_type: constraint
priority: must
verification: inspection of websocket.go for bulk read handlers
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: Bulk fetches belong on REST to keep the socket narrow and cache-friendly.
verified_against:
    websocket-server-hdup:
        hash: c2d6b7c8e035720d03d5754cdb196199ca2da6a30d1db538fa95a0ba7fe31fc5
        at: "2026-04-18T09:37:57Z"
---
