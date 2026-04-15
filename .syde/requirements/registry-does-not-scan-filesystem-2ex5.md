---
id: REQ-0067
kind: requirement
name: Registry Does Not Scan Filesystem
slug: registry-does-not-scan-filesystem-2ex5
relationships:
    - target: project-registry-q1zx
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:41Z"
statement: The project registry shall not scan the filesystem for syde projects.
req_type: constraint
priority: must
verification: inspection of registry.go for filesystem walks
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: Implicit discovery would surprise users and conflict with explicit CLI registration.
---
