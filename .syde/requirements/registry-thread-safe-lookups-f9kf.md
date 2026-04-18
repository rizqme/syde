---
id: REQ-0065
kind: requirement
name: Registry Thread Safe Lookups
slug: registry-thread-safe-lookups-f9kf
relationships:
    - target: project-registry-q1zx
      type: refines
updated_at: "2026-04-18T09:37:52Z"
statement: The project registry shall provide thread-safe lookups when accessed concurrently by HTTP handlers.
req_type: non-functional
priority: must
verification: race-detector integration test with parallel GetStore calls
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: HTTP handlers serve requests concurrently and must share one cache safely.
verified_against:
    project-registry-q1zx:
        hash: fc59bad1a4ec2326e7c2fe30774517a5d8c6f333be0f6b67b43ec27c34562297
        at: "2026-04-18T09:37:52Z"
---
