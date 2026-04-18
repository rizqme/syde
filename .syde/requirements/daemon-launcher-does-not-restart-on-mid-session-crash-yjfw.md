---
id: REQ-0136
kind: requirement
name: Daemon Launcher Does Not Restart On Mid Session Crash
slug: daemon-launcher-does-not-restart-on-mid-session-crash-yjfw
relationships:
    - target: daemon-launcher-tzso
      type: refines
updated_at: "2026-04-18T09:37:16Z"
statement: The daemon launcher shall not restart syded if it crashes after the initial readiness handshake succeeds.
req_type: constraint
priority: must
verification: code review for post-start restart handling
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Mid-session recovery is deliberately out of scope to keep launcher logic simple.
verified_against:
    daemon-launcher-tzso:
        hash: 6eff903308820d498bcecf7735f691f70ee70cb5fd4b94866e01c29ffa0a9645
        at: "2026-04-18T09:37:16Z"
---
