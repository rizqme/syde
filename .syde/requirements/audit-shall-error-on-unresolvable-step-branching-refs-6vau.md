---
id: REQ-0354
kind: requirement
name: Audit shall error on unresolvable step branching refs
slug: audit-shall-error-on-unresolvable-step-branching-refs-6vau
description: ERROR when on_success/on_failure references nonexistent step ID
relationships:
    - target: syde
      type: belongs_to
    - target: audit-engine
      type: refines
updated_at: "2026-04-16T10:40:59Z"
statement: If a flow step on_success or on_failure field references a step ID that does not exist in the same flow, then the syde audit engine shall report an error.
req_type: functional
priority: must
verification: Nonexistent step ID causes syde sync check to error
source: plan
requirement_status: active
rationale: Graph integrity is a precondition for correct flowchart rendering
---
