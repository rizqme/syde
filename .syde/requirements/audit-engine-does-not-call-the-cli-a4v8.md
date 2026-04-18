---
id: REQ-0058
kind: requirement
name: Audit Engine Does Not Call the CLI
slug: audit-engine-does-not-call-the-cli-a4v8
relationships:
    - target: audit-engine-4ktg
      type: refines
updated_at: "2026-04-18T09:37:54Z"
statement: The audit engine shall not invoke the syde CLI or any shell command during an audit run.
req_type: constraint
priority: must
verification: code review of internal/audit imports
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: The audit engine is a library used by the CLI, not a consumer of it, to avoid recursion and lock contention.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:37:54Z"
---
