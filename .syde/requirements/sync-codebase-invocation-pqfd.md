---
id: REQ-0307
kind: requirement
name: Sync Codebase Invocation
slug: sync-codebase-invocation-pqfd
relationships:
    - target: sync-codebase-b4jw
      type: refines
    - target: audit-engine-4ktg
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:42Z"
statement: When the user runs syde sync, the syde CLI shall print a sync and audit report highlighting missing components, files, and relationships on stdout.
req_type: interface
priority: must
verification: integration test invoking syde sync
source: manual
source_ref: contract:sync-codebase-b4jw
requirement_status: active
rationale: Sync is the primary gap-audit command that keeps the model aligned with the codebase.
verified_against:
    audit-engine-4ktg:
        hash: ef935d985e980f0e4d1d7dd4eab9ff52c9d0581e2f5f4a976302a3ba33822fbe
        at: "2026-04-18T09:36:42Z"
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:42Z"
---
