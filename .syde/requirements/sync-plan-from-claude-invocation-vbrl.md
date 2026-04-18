---
id: REQ-0310
kind: requirement
name: Sync Plan From Claude Invocation
slug: sync-plan-from-claude-invocation-vbrl
relationships:
    - target: sync-plan-from-claude-rgaw
      type: refines
    - target: cli-commands-hpjb
      type: refines
updated_at: "2026-04-18T09:37:38Z"
statement: When the user runs syde plan sync, the syde CLI shall import the plan from a .claude/plans markdown file into the syde model and print the new plan slug.
req_type: interface
priority: must
verification: integration test invoking syde plan sync against a sample .claude plan file
source: manual
source_ref: contract:sync-plan-from-claude-rgaw
requirement_status: active
rationale: Plan sync bridges Claude Code-authored plans into the syde governance model.
verified_against:
    cli-commands-hpjb:
        hash: 9ac9799d70204c6b6eb6e65813516e5b489e49982b2d96ec4becb22a2d4dfbc5
        at: "2026-04-18T09:37:38Z"
---
