---
id: REQ-0206
kind: requirement
name: Project Overview Returns Overview JSON
slug: project-overview-returns-overview-json-rwl6
relationships:
    - target: project-overview-j6y9
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:13Z"
statement: When a client invokes GET /api/<project>/overview, the syded daemon shall respond with 200 OK and a JSON body containing name, entity_counts, and recent_changes.
req_type: interface
priority: must
verification: integration test against /api/<project>/overview
source: manual
source_ref: contract:project-overview-j6y9
requirement_status: active
rationale: The dashboard overview page needs aggregate project stats.
---
