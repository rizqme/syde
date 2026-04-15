---
id: REQ-0058
kind: requirement
name: Audit Engine Does Not Call the CLI
slug: audit-engine-does-not-call-the-cli-a4v8
relationships:
    - target: audit-engine-4ktg
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:35Z"
statement: The audit engine shall not invoke the syde CLI or any shell command during an audit run.
req_type: constraint
priority: must
verification: code review of internal/audit imports
source: manual
source_ref: component:audit-engine-4ktg
requirement_status: active
rationale: The audit engine is a library used by the CLI, not a consumer of it, to avoid recursion and lock contention.
---
