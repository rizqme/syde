---
id: REQ-0076
kind: requirement
name: CLI Commands Support Summary Tree Management
slug: cli-commands-support-summary-tree-management-d0so
relationships:
    - target: cli-commands-hpjb
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:52:51Z"
statement: The syde CLI shall provide tree scan, summarize, show, context, and status subcommands for managing the summary tree.
req_type: functional
priority: must
verification: integration test invoking syde tree subcommands
source: manual
source_ref: component:cli-commands-hpjb
requirement_status: active
rationale: Summary tree maintenance is the entry point to Phase 0 bootstrapping.
---
