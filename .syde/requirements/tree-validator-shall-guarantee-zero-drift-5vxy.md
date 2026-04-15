---
id: REQ-0009
kind: requirement
name: Tree validator shall guarantee zero drift
slug: tree-validator-shall-guarantee-zero-drift-5vxy
relationships:
    - target: syde-5tdt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:49:02Z"
statement: The syde validator shall report any source file whose hash changed after the owning entity was last updated.
req_type: non-functional
priority: must
verification: syde sync check --strict on a repo with a touched file emits a drift WARN
source: manual
source_ref: system:syde-5tdt:quality_goals
requirement_status: active
rationale: Drift detection is the only mechanism keeping the design model in lock-step with the code.
---
