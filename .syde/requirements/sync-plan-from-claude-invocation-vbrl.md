---
id: REQ-0310
kind: requirement
name: Sync Plan From Claude Invocation
slug: sync-plan-from-claude-invocation-vbrl
relationships:
    - target: sync-plan-from-claude-rgaw
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:02:37Z"
statement: When the user runs syde plan sync, the syde CLI shall import the plan from a .claude/plans markdown file into the syde model and print the new plan slug.
req_type: interface
priority: must
verification: integration test invoking syde plan sync against a sample .claude plan file
source: manual
source_ref: contract:sync-plan-from-claude-rgaw
requirement_status: active
rationale: Plan sync bridges Claude Code-authored plans into the syde governance model.
---
