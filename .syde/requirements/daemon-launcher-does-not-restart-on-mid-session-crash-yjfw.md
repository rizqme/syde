---
id: REQ-0136
kind: requirement
name: Daemon Launcher Does Not Restart On Mid Session Crash
slug: daemon-launcher-does-not-restart-on-mid-session-crash-yjfw
relationships:
    - target: daemon-launcher-tzso
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:08Z"
statement: The daemon launcher shall not restart syded if it crashes after the initial readiness handshake succeeds.
req_type: constraint
priority: must
verification: code review for post-start restart handling
source: manual
source_ref: component:daemon-launcher-tzso
requirement_status: active
rationale: Mid-session recovery is deliberately out of scope to keep launcher logic simple.
---
