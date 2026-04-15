---
id: REQ-0239
kind: requirement
name: Graph Depth Parameter Controls Traversal
slug: graph-depth-parameter-controls-traversal-v4k2
relationships:
    - target: graph-screen-4044
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:54Z"
statement: Where a depth query parameter is provided on the /__graph__ route, the dashboard shall traverse the relationship graph to that depth, defaulting to two when absent.
req_type: interface
priority: should
verification: manual inspection of /__graph__?depth=3 in the dashboard
source: manual
source_ref: contract:graph-screen-4044
requirement_status: active
rationale: Depth control lets users widen or narrow the relationship neighborhood.
---
