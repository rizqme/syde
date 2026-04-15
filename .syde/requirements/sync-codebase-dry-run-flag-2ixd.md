---
id: REQ-0309
kind: requirement
name: Sync Codebase Dry Run Flag
slug: sync-codebase-dry-run-flag-2ixd
relationships:
    - target: sync-codebase-b4jw
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: Where --dry-run is passed to syde sync, the syde CLI shall report gaps without writing any changes to disk.
req_type: interface
priority: must
verification: integration test invoking syde sync --dry-run and checking the model was unmodified
source: manual
source_ref: contract:sync-codebase-b4jw
requirement_status: active
rationale: Dry-run support lets operators preview sync effects before committing them.
---
