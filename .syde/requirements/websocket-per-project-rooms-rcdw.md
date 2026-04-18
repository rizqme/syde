---
id: REQ-0102
kind: requirement
name: WebSocket Per Project Rooms
slug: websocket-per-project-rooms-rcdw
relationships:
    - target: websocket-server-hdup
      type: refines
updated_at: "2026-04-18T09:38:11Z"
statement: The syded WebSocket server shall isolate broadcasts to per-project rooms so events from one project are not delivered to another.
req_type: functional
priority: must
verification: end-to-end test connecting two clients to different projects and asserting no cross-delivery
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: Cross-project leakage would surprise users running multiple projects on one daemon.
verified_against:
    websocket-server-hdup:
        hash: c2d6b7c8e035720d03d5754cdb196199ca2da6a30d1db538fa95a0ba7fe31fc5
        at: "2026-04-18T09:38:11Z"
---
