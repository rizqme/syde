---
id: REQ-0488
kind: requirement
name: Dashboard graph shall render every system at the same visual tier
slug: dashboard-graph-shall-render-every-system-at-the-same-visual-tier-pb2g
relationships:
    - target: web-spa-jy9z
      type: refines
updated_at: "2026-04-18T09:55:14Z"
statement: The syded dashboard graph shall render every system entity using a single size and colour without distinguishing root from sub-systems.
req_type: functional
priority: should
verification: Loading the /graph route shows every system entity at the same node radius with a single legend entry
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Mirrors the flat system topology in the model.
verified_against:
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:55:14Z"
---
