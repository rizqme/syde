---
id: REQ-0239
kind: requirement
name: Graph Depth Parameter Controls Traversal
slug: graph-depth-parameter-controls-traversal-v4k2
relationships:
    - target: graph-screen-4044
      type: refines
    - target: web-spa-jy9z
      type: refines
    - target: graph-engine-xgjy
      type: refines
updated_at: "2026-04-18T09:36:46Z"
statement: Where a depth query parameter is provided on the /__graph__ route, the dashboard shall traverse the relationship graph to that depth, defaulting to two when absent.
req_type: interface
priority: should
verification: manual inspection of /__graph__?depth=3 in the dashboard
source: manual
source_ref: contract:graph-screen-4044
requirement_status: active
rationale: Depth control lets users widen or narrow the relationship neighborhood.
verified_against:
    graph-engine-xgjy:
        hash: 008188a7a397c93a8d847fa561e5274e77480780d36faacff440a814f6d605fe
        at: "2026-04-18T09:36:46Z"
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:36:46Z"
---
