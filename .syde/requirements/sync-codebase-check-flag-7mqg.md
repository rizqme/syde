---
id: REQ-0308
kind: requirement
name: Sync Codebase Check Flag
slug: sync-codebase-check-flag-7mqg
relationships:
    - target: sync-codebase-b4jw
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:32Z"
statement: Where --check is passed to syde sync, the syde CLI shall exit with a non-zero status code whenever completeness gaps are detected.
req_type: interface
priority: must
verification: integration test invoking syde sync --check against a project with known gaps
source: manual
source_ref: contract:sync-codebase-b4jw
requirement_status: active
rationale: Non-zero exit on gaps enables CI enforcement of design-code alignment.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:32Z"
---
