---
id: REQ-0010
kind: requirement
name: CLI commands shall be idempotent
slug: cli-commands-shall-be-idempotent-dvyb
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:38:03Z"
statement: The syde CLI shall produce the same end state when any command is re-run with identical arguments.
req_type: non-functional
priority: must
verification: Smoke test running every write command twice in succession
source: manual
source_ref: system:syde-5tdt:quality_goals
requirement_status: active
rationale: Agents retry on failure; non-idempotent commands cause duplicate entities and corrupted state.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:38:03Z"
---
