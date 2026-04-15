---
id: REQ-0203
kind: requirement
name: List Projects Returns Projects Array
slug: list-projects-returns-projects-array-bc1r
relationships:
    - target: list-projects-eilx
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:05Z"
statement: When a client invokes GET /api/projects, the syded daemon shall respond with 200 OK and a JSON body containing a projects array of registered project entries.
req_type: interface
priority: must
verification: integration test against /api/projects
source: manual
source_ref: contract:list-projects-eilx
requirement_status: active
rationale: The dashboard SPA needs to list registered projects on load.
---
