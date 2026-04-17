---
id: REQ-0358
kind: requirement
name: Flowchart success edges shall be solid green
slug: flowchart-success-edges-shall-be-solid-green-ow1t
description: on_success edges render as solid green arrows
relationships:
    - target: syde
      type: belongs_to
    - target: web-spa
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: The dashboard flowchart shall render on_success edges as solid green arrows.
req_type: usability
priority: should
verification: Flowchart shows green solid arrows for success paths
source: plan
requirement_status: active
rationale: Color-coded edges make branching immediately scannable
---
