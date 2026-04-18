---
id: REQ-0497
kind: requirement
name: 'Approved plan: Remove root system and allow components to belong to multiple standal...'
slug: approved-plan-remove-root-system-and-allow-components-to-belong-to-multiple-standal-5cvq
relationships:
    - target: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
      type: references
      label: approved_plan
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T10:00:44Z"
statement: When PLN-0019 is approved, the syde design model shall reflect a flat two-system topology with no root system and bidirectional requirement-component coverage.
req_type: functional
priority: must
verification: syde sync check exits 0 and syde list system returns exactly syde and syded
source: plan
source_ref: plan:remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Approval requirement auto-created by syde plan approve for PLN-0019
approved_at: "2026-04-18T10:00:13Z"
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T10:00:44Z"
---
