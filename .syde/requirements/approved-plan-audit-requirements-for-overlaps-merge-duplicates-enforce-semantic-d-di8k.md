---
id: REQ-0394
kind: requirement
name: 'Approved plan: Audit requirements for overlaps, merge duplicates, enforce semantic d...'
slug: approved-plan-audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-d-di8k
relationships:
    - target: audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
      type: references
      label: approved_plan
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T11:03:29Z"
statement: The syde audit engine shall enforce semantic-distinction acknowledgement on every requirement overlap, contract co-evolution on every requirement naming a CLI REST screen or event surface, and a post-plan counterpart for every planning-time rule.
req_type: constraint
priority: must
verification: sync check errors on rubber-stamp acknowledgements, missing contract coverage, and any plan_authoring rule without a post-plan twin
source: plan
source_ref: plan:audit-requirements-for-overlaps-merge-duplicates-enforce-semantic-distinction-at-the-harness-level-rguz
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-17T10:02:42Z"
---
