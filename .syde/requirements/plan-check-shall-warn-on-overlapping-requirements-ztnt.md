---
id: REQ-0344
kind: requirement
name: Plan check shall warn on overlapping requirements
slug: plan-check-shall-warn-on-overlapping-requirements-ztnt
relationships:
    - target: syde
      type: belongs_to
    - target: audit-engine
      type: refines
updated_at: "2026-04-16T09:48:49Z"
statement: When a new requirement in a plan shares significant terms with an existing active requirement, the syde audit engine shall warn that the requirements may overlap.
req_type: functional
priority: must
verification: A plan adding a requirement similar to an existing one triggers the overlap warning
source: plan
requirement_status: active
rationale: Overlapping requirements should be linked via refines/derives_from or the old one superseded
---
