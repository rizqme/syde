---
id: REQ-0008
kind: requirement
name: Architecture context shall load fast
slug: architecture-context-shall-load-fast-52uc
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:35Z"
statement: When a session starts, the syde context bridge shall load and inject the full architecture snapshot in under one second for projects with fewer than five hundred entities.
req_type: performance
priority: must
verification: syde context --json elapsed time measured during smoke test
source: manual
source_ref: system:syde-5tdt:quality_goals
requirement_status: active
rationale: Slow startup defeats the auto-load workflow that makes syde useful to agents.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:35Z"
---
