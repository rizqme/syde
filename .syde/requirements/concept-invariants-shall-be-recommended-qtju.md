---
id: REQ-0372
kind: requirement
name: Concept invariants shall be recommended
slug: concept-invariants-shall-be-recommended-qtju
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T08:25:33Z"
statement: The syde audit engine shall warn when a concept entity has an empty invariants field.
req_type: functional
priority: should
verification: sync check warns on empty invariants
source: plan
requirement_status: active
rationale: Invariants document rules
---
