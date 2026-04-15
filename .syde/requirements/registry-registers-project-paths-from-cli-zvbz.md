---
id: REQ-0063
kind: requirement
name: Registry Registers Project Paths From CLI
slug: registry-registers-project-paths-from-cli-zvbz
relationships:
    - target: project-registry-q1zx
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:52:41Z"
statement: When a CLI invocation references a new project path, the project registry shall register that path for subsequent lookups.
req_type: functional
priority: must
verification: integration test posting a new project path and then reading it back
source: manual
source_ref: component:project-registry-q1zx
requirement_status: active
rationale: CLI invocations are the source of truth for which projects the daemon knows.
---
