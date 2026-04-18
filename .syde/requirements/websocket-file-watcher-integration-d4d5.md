---
id: REQ-0103
kind: requirement
name: WebSocket File Watcher Integration
slug: websocket-file-watcher-integration-d4d5
relationships:
    - target: websocket-server-hdup
      type: refines
updated_at: "2026-04-18T09:37:07Z"
statement: When a file watcher detects a change inside a project's .syde directory, the syded WebSocket server shall publish a refresh event to that project's room.
req_type: functional
priority: must
verification: end-to-end test writing to a project file and asserting a websocket frame
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: Disk changes are the trigger for every live refresh.
verified_against:
    websocket-server-hdup:
        hash: c2d6b7c8e035720d03d5754cdb196199ca2da6a30d1db538fa95a0ba7fe31fc5
        at: "2026-04-18T09:37:07Z"
---
