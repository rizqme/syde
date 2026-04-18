---
id: REQ-0246
kind: requirement
name: Get Tree Node Returns Breadcrumb Ancestors
slug: get-tree-node-returns-breadcrumb-ancestors-f12h
relationships:
    - target: get-tree-node-http-uqzq
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:43Z"
statement: When GET /api/<project>/tree/<path> succeeds, the syded daemon shall return breadcrumb as an array of ancestor folder summary objects.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree/<path>
source: manual
source_ref: contract:get-tree-node-http-uqzq
requirement_status: active
rationale: The tree node view renders an ancestor path breadcrumb.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:43Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:43Z"
---
