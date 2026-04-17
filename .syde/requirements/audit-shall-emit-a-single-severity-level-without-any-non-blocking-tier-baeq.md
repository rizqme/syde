---
id: REQ-0404
kind: requirement
name: Audit shall emit a single severity level without any non-blocking tier
slug: audit-shall-emit-a-single-severity-level-without-any-non-blocking-tier-baeq
relationships:
    - target: audit-engine
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-17T10:52:32Z"
statement: The syde audit engine shall emit findings at a single blocking severity level, without any warning or hint tier, so that every finding causes the sync-check, validate, and plan-complete gates to exit non-zero.
req_type: constraint
priority: must
verification: grep for SeverityWarning or SeverityHint in internal/audit returns zero results and a Go test asserts the audit severity enum has exactly one exported value
source: plan
requirement_status: active
rationale: A strict project cannot afford a non-blocking audit tier; findings either matter or they should not be emitted.
---
