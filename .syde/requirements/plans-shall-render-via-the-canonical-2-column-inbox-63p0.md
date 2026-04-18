---
id: REQ-0334
kind: requirement
name: Plans shall render via the canonical 2-column inbox
slug: plans-shall-render-via-the-canonical-2-column-inbox-63p0
description: Plans inbox layout requirement for the syded dashboard.
relationships:
    - target: web-spa-jy9z
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:39Z"
statement: The syded dashboard shall render the Plans inbox as a 2-column layout with the entity list on the left and the selected plan's detail rendered inline in the right column.
req_type: usability
priority: must
verification: 'Manual inspection of /<project>/plan in the dashboard: a plan list is on the left, selecting a plan renders the detail inline on the right, no floating panel appears.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: Plans need the same inbox UX as every other entity kind so reviewers do not learn a separate navigation pattern, and the floating EntityDetail panel becomes dead weight.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:39Z"
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:39Z"
---
