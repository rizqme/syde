---
id: REQ-0356
kind: requirement
name: Audit shall warn on steps with empty contract
slug: audit-shall-warn-on-steps-with-empty-contract-fvvs
description: WARN when step has no contract ref
relationships:
- target: audit-engine
  type: refines
updated_at: '2026-04-17T11:05:30Z'
statement: When a flow step has an empty contract field, the syde audit engine shall report a warning.
req_type: functional
priority: should
verification: Step without contract slug causes syde sync check to warn
source: plan
requirement_status: obsolete
rationale: Internal steps are legitimate but should be visible for review
obsolete_reason: Internal steps without a contract ref are legitimate per the strict-audit redesign — the audit no longer warns on empty contract fields. A genuinely missing contract is caught instead by the unresolvable-contract-ref rule.
---
