---
id: REQ-0065
kind: requirement
name: Registry Thread Safe Lookups
slug: registry-thread-safe-lookups-f9kf
relationships:
    - target: project-registry-q1zx
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:41Z"
statement: The project registry shall provide thread-safe lookups when accessed concurrently by HTTP handlers.
req_type: non-functional
priority: must
verification: race-detector integration test with parallel GetStore calls
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: HTTP handlers serve requests concurrently and must share one cache safely.
---
