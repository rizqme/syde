---
id: REQ-0403
kind: requirement
name: Plan diffs introducing contracts shall also introduce matching flow steps
slug: plan-diffs-introducing-contracts-shall-also-introduce-matching-flow-steps-s8at
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:32Z"
statement: When a syde plan change diff adds or extends a contract entry, the same diff shall pair that entry with a flow lane item whose trajectory names the contract slug.
req_type: functional
priority: must
verification: plan check on a diff that adds a contract without a paired flow lane item emits a warning; pairing clears it
source: plan
requirement_status: active
rationale: Enforces co-evolution of boundary definitions and their traversal documents at planning time.
---
