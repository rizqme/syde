---
id: REQ-0338
kind: requirement
name: Plan changes view shall group Extended by target
slug: plan-changes-view-shall-group-extended-by-target-6qld
description: Plan changes grouping requirement for Extended entries with the same target.
relationships:
    - target: web-spa-jy9z
      type: refines
    - target: http-api-afos
      type: refines
updated_at: "2026-04-18T09:37:11Z"
statement: When the plan changes view renders multiple Extended changes targeting the same entity slug, the syded dashboard shall group them into a single card with all what, why, and field changes stacked inside.
req_type: usability
priority: must
verification: 'Manual inspection: open a plan with two Extended changes on the same target and confirm only one card renders for that target.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: Reviewers think in terms of what is changing on an entity, not which change entries exist. The card boundary should match the entity boundary.
verified_against:
    http-api-afos:
        hash: ab080a2b2498114076ebb7cb0bdfeb444f53e7a3af2f5af4bd111c0b11855b65
        at: "2026-04-18T09:37:11Z"
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:11Z"
---
