---
id: REQ-0008
kind: requirement
name: Architecture context shall load fast
slug: architecture-context-shall-load-fast-52uc
relationships:
    - target: syde-5tdt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:49:02Z"
statement: When a session starts, the syde context bridge shall load and inject the full architecture snapshot in under one second for projects with fewer than five hundred entities.
req_type: performance
priority: must
verification: syde context --json elapsed time measured during smoke test
source: manual
source_ref: system:syde-5tdt:quality_goals
requirement_status: active
rationale: Slow startup defeats the auto-load workflow that makes syde useful to agents.
---
