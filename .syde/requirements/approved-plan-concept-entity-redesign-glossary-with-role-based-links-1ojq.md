---
id: REQ-0367
kind: requirement
name: 'Approved plan: Concept entity redesign: glossary with role-based links'
slug: approved-plan-concept-entity-redesign-glossary-with-role-based-links-1ojq
relationships:
    - target: concept-entity-redesign-glossary-with-role-based-links-uxm1
      type: references
      label: approved_plan
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:19:28Z"
statement: The syde concept entity shall be a glossary term whose only fields are meaning invariants and lifecycle, with role-based relationships implemented_by exposed_via and used_in replacing the former ERD canvas.
req_type: functional
priority: must
verification: ConceptEntity carries only the three glossary fields, the dashboard renders the 2-column glossary layout, and the new relationship types are valid
source: plan
source_ref: plan:concept-entity-redesign-glossary-with-role-based-links-uxm1
requirement_status: active
rationale: Captured automatically when the plan was approved.
approved_at: "2026-04-17T02:45:17Z"
---
