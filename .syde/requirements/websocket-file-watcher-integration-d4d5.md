---
id: REQ-0103
kind: requirement
name: WebSocket File Watcher Integration
slug: websocket-file-watcher-integration-d4d5
relationships:
    - target: websocket-server-hdup
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:53:16Z"
statement: When a file watcher detects a change inside a project's .syde directory, the syded WebSocket server shall publish a refresh event to that project's room.
req_type: functional
priority: must
verification: end-to-end test writing to a project file and asserting a websocket frame
source: manual
source_ref: component:websocket-server-hdup
requirement_status: active
rationale: Disk changes are the trigger for every live refresh.
---
