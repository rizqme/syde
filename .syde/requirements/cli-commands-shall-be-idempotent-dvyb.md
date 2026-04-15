---
id: REQ-0010
kind: requirement
name: CLI commands shall be idempotent
slug: cli-commands-shall-be-idempotent-dvyb
relationships:
    - target: syde-5tdt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:49:02Z"
statement: The syde CLI shall produce the same end state when any command is re-run with identical arguments.
req_type: non-functional
priority: must
verification: Smoke test running every write command twice in succession
source: manual
source_ref: system:syde-5tdt:quality_goals
requirement_status: active
rationale: Agents retry on failure; non-idempotent commands cause duplicate entities and corrupted state.
---
