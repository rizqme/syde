---
id: REQ-0203
kind: requirement
name: List Projects Returns Projects Array
slug: list-projects-returns-projects-array-bc1r
relationships:
    - target: list-projects-eilx
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:49Z"
statement: When a client invokes GET /api/projects, the syded daemon shall respond with 200 OK and a JSON body containing a projects array of registered project entries.
req_type: interface
priority: must
verification: integration test against /api/projects
source: manual
source_ref: contract:list-projects-eilx
requirement_status: active
rationale: The dashboard SPA needs to list registered projects on load.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:49Z"
---
