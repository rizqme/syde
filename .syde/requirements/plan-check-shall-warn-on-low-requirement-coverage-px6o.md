---
id: REQ-0343
kind: requirement
name: Plan check shall warn on low requirement coverage
slug: plan-check-shall-warn-on-low-requirement-coverage-px6o
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:38Z"
statement: When a plan has fewer requirement changes than one third of its non-requirement changes, the syde audit engine shall warn that the plan may be under-specified.
req_type: functional
priority: must
verification: A plan with 9 component changes and 1 requirement triggers the warning
source: plan
requirement_status: active
rationale: Every design decision should trace to a requirement
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:38Z"
---
