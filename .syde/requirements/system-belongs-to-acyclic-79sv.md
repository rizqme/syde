---
id: REQ-0080
kind: requirement
name: System Belongs-To Acyclic
slug: system-belongs-to-acyclic-79sv
relationships:
    - target: relationship-hjgt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:58Z"
statement: If adding a belongs_to relationship between systems would create a cycle, then the syde CLI shall reject the save with a cycle error.
req_type: constraint
priority: must
verification: integration test adding a belongs_to link that closes a cycle
source: manual
source_ref: concept:relationship-hjgt
requirement_status: active
rationale: Cyclic system nesting breaks ownership trees used by belongs_to queries.
---
