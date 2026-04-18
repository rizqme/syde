---
id: REQ-0309
kind: requirement
name: Sync Codebase Dry Run Flag
slug: sync-codebase-dry-run-flag-2ixd
relationships:
    - target: sync-codebase-b4jw
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:06Z"
statement: Where --dry-run is passed to syde sync, the syde CLI shall report gaps without writing any changes to disk.
req_type: interface
priority: must
verification: integration test invoking syde sync --dry-run and checking the model was unmodified
source: manual
source_ref: contract:sync-codebase-b4jw
requirement_status: active
rationale: Dry-run support lets operators preview sync effects before committing them.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:06Z"
---
