---
id: REQ-0170
kind: requirement
name: Skill Installer Does Not Execute Hooks
slug: skill-installer-does-not-execute-hooks-rj1w
relationships:
    - target: skill-installer-wbmu
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:55:49Z"
statement: The skill installer shall not execute hook scripts and shall only write the hook JSON for Claude Code to consume.
req_type: constraint
priority: must
verification: code review of internal/skill for hook execution paths
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Hook execution is the agent harness's responsibility, not the installer's.
---
