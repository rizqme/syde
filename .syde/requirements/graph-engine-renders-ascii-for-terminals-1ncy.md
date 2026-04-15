---
id: REQ-0149
kind: requirement
name: Graph Engine Renders ASCII For Terminals
slug: graph-engine-renders-ascii-for-terminals-1ncy
relationships:
    - target: graph-engine-xgjy
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:54:49Z"
statement: When RenderASCII is called with a SubgraphResult, the graph engine shall produce a human-readable outbound and inbound edge listing suitable for terminal display.
req_type: functional
priority: must
verification: unit test of RenderASCII in internal/graph/render.go
source: manual
source_ref: component:graph-engine-xgjy
requirement_status: active
rationale: ASCII output is the default format for the syde graph command.
---
