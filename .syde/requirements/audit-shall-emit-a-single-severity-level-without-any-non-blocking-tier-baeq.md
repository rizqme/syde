---
id: REQ-0404
kind: requirement
name: Audit shall emit a single severity level without any non-blocking tier
slug: audit-shall-emit-a-single-severity-level-without-any-non-blocking-tier-baeq
relationships:
    - target: audit-engine
      type: refines
updated_at: "2026-04-18T09:37:13Z"
statement: The syde audit engine shall emit findings at a single blocking severity level, without any warning or hint tier, so that every finding causes the sync-check, validate, and plan-complete gates to exit non-zero.
req_type: constraint
priority: must
verification: grep for SeverityWarning or SeverityHint in internal/audit returns zero results and a Go test asserts the audit severity enum has exactly one exported value
source: plan
requirement_status: active
rationale: A strict project cannot afford a non-blocking audit tier; findings either matter or they should not be emitted.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:13Z"
---
