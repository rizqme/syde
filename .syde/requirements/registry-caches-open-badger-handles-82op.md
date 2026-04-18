---
id: REQ-0064
kind: requirement
name: Registry Caches Open Badger Handles
slug: registry-caches-open-badger-handles-82op
relationships:
    - target: project-registry-q1zx
      type: refines
updated_at: "2026-04-18T09:37:01Z"
statement: The project registry shall cache open BadgerDB handles keyed by project .syde directory for the daemon lifetime.
req_type: functional
priority: must
verification: integration test asserting handle reuse across requests
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: Reopening BadgerDB on every request is both slow and conflicts with single-writer semantics.
verified_against:
    project-registry-q1zx:
        hash: fc59bad1a4ec2326e7c2fe30774517a5d8c6f333be0f6b67b43ec27c34562297
        at: "2026-04-18T09:37:01Z"
---
