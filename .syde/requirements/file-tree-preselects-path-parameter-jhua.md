---
id: REQ-0236
kind: requirement
name: File Tree Preselects Path Parameter
slug: file-tree-preselects-path-parameter-jhua
relationships:
    - target: file-tree-screen-fez2
      type: refines
    - target: web-spa-jy9z
      type: refines
    - target: summary-tree-fq6u
      type: refines
updated_at: "2026-04-18T09:37:32Z"
statement: Where a path query parameter is provided on the /__tree__ route, the dashboard shall preselect the tree node at that path and display its summary and content.
req_type: interface
priority: should
verification: manual inspection of /__tree__?path=... in the dashboard
source: manual
source_ref: contract:file-tree-screen-fez2
requirement_status: active
rationale: Deep linking to a specific file avoids manual tree traversal.
verified_against:
    summary-tree-fq6u:
        hash: 51703195026629fb17ef88e0859de7cdd45e6cd90f54ba62f52398aaf2cb378a
        at: "2026-04-18T09:37:32Z"
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:32Z"
---
