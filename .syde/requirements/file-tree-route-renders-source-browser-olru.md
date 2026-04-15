---
id: REQ-0234
kind: requirement
name: File Tree Route Renders Source Browser
slug: file-tree-route-renders-source-browser-olru
relationships:
    - target: file-tree-screen-fez2
      type: refines
    - target: syded-dashboard-e82c
      type: belongs_to
updated_at: "2026-04-15T10:59:47Z"
statement: When the user navigates to the /__tree__ route, the dashboard shall render a source tree sidebar and a main pane showing the selected file's summary and content.
req_type: interface
priority: must
verification: manual inspection of /__tree__ in the dashboard
source: manual
source_ref: contract:file-tree-screen-fez2
requirement_status: active
rationale: File tree screen is the browsing surface for summarized source files.
---
