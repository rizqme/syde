---
id: REQ-0486
kind: requirement
name: Design model shall not designate a root system
slug: design-model-shall-not-designate-a-root-system-8joh
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:55:15Z"
statement: The syde audit engine shall not require exactly one system entity to be designated as the root.
req_type: functional
priority: must
verification: syde sync check passes with zero, one, or many system entities without belongs_to edges
source: plan
source_ref: remove-root-system-and-allow-components-to-belong-to-multiple-standalone-systems-gtmh
requirement_status: active
rationale: Flat system set reflects actual process boundaries.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:55:15Z"
---
