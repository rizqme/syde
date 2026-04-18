---
id: REQ-0009
kind: requirement
name: Tree validator shall guarantee zero drift
slug: tree-validator-shall-guarantee-zero-drift-5vxy
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:38:02Z"
statement: The syde validator shall report any source file whose hash changed after the owning entity was last updated.
req_type: non-functional
priority: must
verification: syde sync check --strict on a repo with a touched file emits a drift WARN
source: manual
source_ref: system:syde-5tdt:quality_goals
requirement_status: active
rationale: Drift detection is the only mechanism keeping the design model in lock-step with the code.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:38:02Z"
---
