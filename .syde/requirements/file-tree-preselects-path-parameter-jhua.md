---
id: REQ-0236
kind: requirement
name: File Tree Preselects Path Parameter
slug: file-tree-preselects-path-parameter-jhua
relationships:
    - target: file-tree-screen-fez2
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:50Z"
statement: Where a path query parameter is provided on the /__tree__ route, the dashboard shall preselect the tree node at that path and display its summary and content.
req_type: interface
priority: should
verification: manual inspection of /__tree__?path=... in the dashboard
source: manual
source_ref: contract:file-tree-screen-fez2
requirement_status: active
rationale: Deep linking to a specific file avoids manual tree traversal.
---
