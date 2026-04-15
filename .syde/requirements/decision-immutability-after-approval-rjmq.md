---
id: REQ-0016
kind: requirement
name: Decision Immutability After Approval
slug: decision-immutability-after-approval-rjmq
relationships:
    - target: decision-m2um
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:51:42Z"
statement: If a decision has been approved, then the syde CLI shall not allow edits to its statement or rationale and shall require a superseding decision.
req_type: constraint
priority: should
verification: manual review; convention enforced via code review
source: manual
source_ref: concept:decision-m2um
requirement_status: active
rationale: Preserves ADR auditability so historical choices remain intact.
---
