---
id: REQ-0250
kind: requirement
name: Get Tree Node Returns Children For Folders
slug: get-tree-node-returns-children-for-folders-jfzz
relationships:
    - target: get-tree-node-http-uqzq
      type: refines
    - target: summary-tree-fq6u
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:36:53Z"
statement: When GET /api/<project>/tree/<path> resolves to a folder, the syded daemon shall return children as an array of direct child node objects.
req_type: interface
priority: must
verification: integration test against /api/<project>/tree/<path>
source: manual
source_ref: contract:get-tree-node-http-uqzq
requirement_status: active
rationale: Folder views expand to show direct children.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:36:53Z"
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:36:53Z"
---
