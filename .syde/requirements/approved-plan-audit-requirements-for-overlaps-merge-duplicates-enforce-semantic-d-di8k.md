---
id: REQ-0394
kind: requirement
name: 'Approved plan: Audit requirements for overlaps, merge duplicates, enforce semantic d...'
slug: approved-plan-audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-d-di8k
relationships:
    - target: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
      type: references
      label: approved_plan
    - target: web-spa-jy9z
      type: refines
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:37:53Z"
statement: The syde audit engine shall enforce semantic-distinction acknowledgement on every requirement overlap, contract co-evolution on every requirement naming a CLI REST screen or event surface, and a post-plan counterpart for every planning-time rule.
req_type: constraint
priority: must
verification: sync check errors on rubber-stamp acknowledgements, missing contract coverage, and any plan_authoring rule without a post-plan twin
source: plan
source_ref: plan:audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-17T10:02:42Z"
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:53Z"
    web-spa-jy9z:
        hash: 3e31271ac2769b109897c09240242676ec33b6a4c31e4e49f30f94ef09dccb45
        at: "2026-04-18T09:37:53Z"
---
