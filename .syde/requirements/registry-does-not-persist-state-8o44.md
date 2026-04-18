---
id: REQ-0066
kind: requirement
name: Registry Does Not Persist State
slug: registry-does-not-persist-state-8o44
relationships:
    - target: project-registry-q1zx
      type: refines
updated_at: "2026-04-18T09:37:03Z"
statement: The project registry shall not persist its state across daemon restarts.
req_type: constraint
priority: must
verification: inspection of registry.go confirming no on-disk state
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: Persistence is out of scope; the CLI re-registers projects on next use.
verified_against:
    project-registry-q1zx:
        hash: fc59bad1a4ec2326e7c2fe30774517a5d8c6f333be0f6b67b43ec27c34562297
        at: "2026-04-18T09:37:03Z"
---
