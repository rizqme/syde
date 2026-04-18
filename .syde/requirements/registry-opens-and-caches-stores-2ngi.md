---
id: REQ-0062
kind: requirement
name: Registry Opens And Caches Stores
slug: registry-opens-and-caches-stores-2ngi
relationships:
    - target: project-registry-q1zx
      type: refines
updated_at: "2026-04-18T09:38:05Z"
statement: The project registry shall open, cache, and serve a Store handle for each project path on demand.
req_type: functional
priority: must
verification: integration test invoking registry.GetStore twice on the same path and asserting a single open
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: Caching Store handles prevents BadgerDB directory-lock contention under concurrent requests.
verified_against:
    project-registry-q1zx:
        hash: fc59bad1a4ec2326e7c2fe30774517a5d8c6f333be0f6b67b43ec27c34562297
        at: "2026-04-18T09:38:05Z"
---
