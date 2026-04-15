---
id: REQ-0278
kind: requirement
name: Install Skill Invocation
slug: install-skill-invocation-2k4w
relationships:
    - target: install-skill-bji4
      type: refines
    - target: syde-cli-2478
      type: belongs_to
updated_at: "2026-04-15T11:00:55Z"
statement: When the user runs syde install-skill, the syde CLI shall write the skill files into the .claude directory and append the mandatory syde rules section to CLAUDE.md.
req_type: interface
priority: must
verification: integration test invoking syde install-skill and checking for .claude/skills/syde and CLAUDE.md rules block
source: manual
source_ref: contract:install-skill-bji4
requirement_status: active
rationale: Skill installation is how syde injects its enforcement behavior into Claude Code sessions.
---
