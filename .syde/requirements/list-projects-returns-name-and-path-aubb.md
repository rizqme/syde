---
id: REQ-0205
kind: requirement
name: List Projects Returns Name And Path
slug: list-projects-returns-name-and-path-aubb
relationships:
    - target: list-projects-eilx
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:18Z"
statement: When GET /api/projects succeeds, the syded daemon shall return each project entry with a name string and an absolute filesystem path string.
req_type: interface
priority: must
verification: integration test against /api/projects
source: manual
source_ref: contract:list-projects-eilx
requirement_status: active
rationale: Clients resolve project routes and display names using these fields.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:18Z"
---
