---
id: REQ-0105
kind: requirement
name: WebSocket Not A Bulk Fetch Channel
slug: websocket-not-a-bulk-fetch-channel-g5tn
relationships:
    - target: websocket-server-hdup
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:16Z"
statement: The syded WebSocket server shall not serve bulk entity fetches and shall defer bulk data retrieval to the HTTP API.
req_type: constraint
priority: must
verification: inspection of websocket.go for bulk read handlers
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: Bulk fetches belong on REST to keep the socket narrow and cache-friendly.
---
