---
id: REQ-0487
kind: requirement
name: Component shall be allowed to belong to multiple systems
slug: component-shall-be-allowed-to-belong-to-multiple-systems-qd6u
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:55:14Z"
statement: Where a component's implementation files are imported by multiple process entrypoints, the syde design model shall allow the component to carry belongs_to edges to each relevant system entity.
req_type: functional
priority: must
verification: Attaching a second belongs_to:system edge to an existing component passes syde sync check
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Matches the reality that most internal/ packages are linked into both cmd/syde and cmd/syded binaries.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:55:14Z"
---
