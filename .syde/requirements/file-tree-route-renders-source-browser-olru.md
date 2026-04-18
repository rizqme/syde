---
id: REQ-0234
kind: requirement
name: File Tree Route Renders Source Browser
slug: file-tree-route-renders-source-browser-olru
relationships:
    - target: file-tree-screen-fez2
      type: refines
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:37:40Z"
statement: When the user navigates to the /__tree__ route, the dashboard shall render a source tree sidebar and a main pane showing the selected file's summary and content.
req_type: interface
priority: must
verification: manual inspection of /__tree__ in the dashboard
source: manual
source_ref: contract:file-tree-screen-fez2
requirement_status: active
rationale: File tree screen is the browsing surface for summarized source files.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:40Z"
---
