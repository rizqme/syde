---
id: REQ-0016
kind: requirement
name: Decision Immutability After Approval
slug: decision-immutability-after-approval-rjmq
relationships:
    - target: decision-m2um
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:50Z"
statement: If a decision has been approved, then the syde CLI shall not allow edits to its statement or rationale and shall require a superseding decision.
req_type: constraint
priority: should
verification: manual review; convention enforced via code review
source: manual
source_ref: concept:decision-m2um
requirement_status: active
rationale: Preserves ADR auditability so historical choices remain intact.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:50Z"
---
