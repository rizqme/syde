---
id: REQ-0067
kind: requirement
name: Registry Does Not Scan Filesystem
slug: registry-does-not-scan-filesystem-2ex5
relationships:
    - target: project-registry-q1zx
      type: refines
updated_at: "2026-04-18T09:36:54Z"
statement: The project registry shall not scan the filesystem for syde projects.
req_type: constraint
priority: must
verification: inspection of registry.go for filesystem walks
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: Implicit discovery would surprise users and conflict with explicit CLI registration.
verified_against:
    project-registry-q1zx:
        hash: fc59bad1a4ec2326e7c2fe30774517a5d8c6f333be0f6b67b43ec27c34562297
        at: "2026-04-18T09:36:54Z"
---
