---
id: REQ-0308
kind: requirement
name: Sync Codebase Check Flag
slug: sync-codebase-check-flag-7mqg
relationships:
    - target: sync-codebase-b4jw
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: Where --check is passed to syde sync, the syde CLI shall exit with a non-zero status code whenever completeness gaps are detected.
req_type: interface
priority: must
verification: integration test invoking syde sync --check against a project with known gaps
source: manual
source_ref: contract:sync-codebase-b4jw
requirement_status: active
rationale: Non-zero exit on gaps enables CI enforcement of design-code alignment.
---
