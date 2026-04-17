---
id: REQ-0338
kind: requirement
name: Plan changes view shall group Extended by target
slug: plan-changes-view-shall-group-extended-by-target-6qld
description: Plan changes grouping requirement for Extended entries with the same target.
relationships:
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-16T01:09:41Z"
statement: When the plan changes view renders multiple Extended changes targeting the same entity slug, the syded dashboard shall group them into a single card with all what, why, and field changes stacked inside.
req_type: usability
priority: must
verification: 'Manual inspection: open a plan with two Extended changes on the same target and confirm only one card renders for that target.'
source: plan
source_ref: plan:plans-inbox-2-column-layout-fud8
requirement_status: active
rationale: Reviewers think in terms of what is changing on an entity, not which change entries exist. The card boundary should match the entity boundary.
---
