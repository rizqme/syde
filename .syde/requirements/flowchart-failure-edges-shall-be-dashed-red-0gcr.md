---
id: REQ-0359
kind: requirement
name: Flowchart failure edges shall be dashed red
slug: flowchart-failure-edges-shall-be-dashed-red-0gcr
description: on_failure edges render as dashed red arrows
relationships:
    - target: syde
      type: belongs_to
    - target: web-spa
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: The dashboard flowchart shall render on_failure edges as dashed red arrows.
req_type: usability
priority: should
verification: Flowchart shows red dashed arrows for failure paths
source: plan
requirement_status: active
rationale: Color-coded edges make branching immediately scannable
---
