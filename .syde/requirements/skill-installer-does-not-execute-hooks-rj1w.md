---
id: REQ-0170
kind: requirement
name: Skill Installer Does Not Execute Hooks
slug: skill-installer-does-not-execute-hooks-rj1w
relationships:
    - target: skill-installer-wbmu
      type: refines
updated_at: "2026-04-18T10:04:49Z"
statement: The skill installer shall not execute hook scripts and shall only write the hook JSON for Claude Code to consume.
req_type: constraint
priority: must
verification: code review of internal/skill for hook execution paths
source: manual
source_ref: component:skill-installer-wbmu
requirement_status: active
rationale: Hook execution is the agent harness's responsibility, not the installer's.
verified_against:
    skill-installer-wbmu:
        hash: cffead9ff459eb538d256d9a782208243779e6c2132e2e5437b9c07de9b37e20
        at: "2026-04-18T10:04:49Z"
---
