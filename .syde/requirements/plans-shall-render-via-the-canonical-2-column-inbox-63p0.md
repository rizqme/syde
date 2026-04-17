---
id: REQ-0334
kind: requirement
name: Plans shall render via the canonical 2-column inbox
slug: plans-shall-render-via-the-canonical-2-column-inbox-63p0
description: Plans inbox layout requirement for the syded dashboard.
relationships:
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-16T01:33:47Z"
statement: The syded dashboard shall render the Plans inbox as a 2-column layout with the entity list on the left and the selected plan's detail rendered inline in the right column.
req_type: usability
priority: must
verification: 'Manual inspection of /<project>/plan in the dashboard: a plan list is on the left, selecting a plan renders the detail inline on the right, no floating panel appears.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: Plans need the same inbox UX as every other entity kind so reviewers do not learn a separate navigation pattern, and the floating EntityDetail panel becomes dead weight.
---
