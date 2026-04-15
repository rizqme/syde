---
id: REQ-0205
kind: requirement
name: List Projects Returns Name And Path
slug: list-projects-returns-name-and-path-aubb
relationships:
    - target: list-projects-eilx
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:11Z"
statement: When GET /api/projects succeeds, the syded daemon shall return each project entry with a name string and an absolute filesystem path string.
req_type: interface
priority: must
verification: integration test against /api/projects
source: manual
source_ref: contract:list-projects-eilx
requirement_status: active
rationale: Clients resolve project routes and display names using these fields.
---
