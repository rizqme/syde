---
id: REQ-0206
kind: requirement
name: Project Overview Returns Overview JSON
slug: project-overview-returns-overview-json-rwl6
relationships:
    - target: project-overview-j6y9
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:26Z"
statement: When a client invokes GET /api/<project>/overview, the syded daemon shall respond with 200 OK and a JSON body containing name, entity_counts, and recent_changes.
req_type: interface
priority: must
verification: integration test against /api/<project>/overview
source: manual
source_ref: contract:project-overview-j6y9
requirement_status: active
rationale: The dashboard overview page needs aggregate project stats.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:26Z"
---
