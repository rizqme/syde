---
id: REQ-0307
kind: requirement
name: Sync Codebase Invocation
slug: sync-codebase-invocation-pqfd
relationships:
    - target: sync-codebase-b4jw
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde sync, the syde CLI shall print a sync and audit report highlighting missing components, files, and relationships on stdout.
req_type: interface
priority: must
verification: integration test invoking syde sync
source: manual
source_ref: contract:sync-codebase-b4jw
requirement_status: active
rationale: Sync is the primary gap-audit command that keeps the model aligned with the codebase.
---
