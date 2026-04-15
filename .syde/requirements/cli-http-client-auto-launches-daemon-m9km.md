---
id: REQ-0099
kind: requirement
name: CLI HTTP Client Auto Launches Daemon
slug: cli-http-client-auto-launches-daemon-m9km
relationships:
    - target: cli-http-client-otp2
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:53:15Z"
statement: When the CLI HTTP client makes its first request, the client shall invoke daemon.EnsureRunning to auto-launch syded if it is not already running.
req_type: functional
priority: must
verification: integration test that first request auto-starts syded
source: manual
source_ref: component:cli-http-client-otp2
requirement_status: active
rationale: Users should never have to manually start syded.
---
