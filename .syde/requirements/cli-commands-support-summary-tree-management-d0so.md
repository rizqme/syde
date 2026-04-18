---
id: REQ-0076
kind: requirement
name: CLI Commands Support Summary Tree Management
slug: cli-commands-support-summary-tree-management-d0so
relationships:
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:36:38Z"
statement: The syde CLI shall provide tree scan, summarize, show, context, and status subcommands for managing the summary tree.
req_type: functional
priority: must
verification: integration test invoking syde tree subcommands
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: Summary tree maintenance is the entry point to Phase 0 bootstrapping.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:36:38Z"
---
