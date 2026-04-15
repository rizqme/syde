---
id: REQ-0012
kind: requirement
name: Workflows shall be enforced by the skill
slug: workflows-shall-be-enforced-by-the-skill-hqqc
relationships:
    - target: syde-5tdt
      type: refines
    - target: syde-5tdt
      type: belongs_to
updated_at: "2026-04-15T10:49:02Z"
statement: The syde install-skill command shall write hooks that block session-end and code modifications when the syde model is out of sync.
req_type: constraint
priority: must
verification: Inspect installed .claude/hooks/syde-hooks.json for SessionStart/Stop/PostToolUse entries
source: manual
source_ref: system:syde-5tdt:design_principles
requirement_status: active
rationale: Without enforcement the model rots; the skill is the only mechanism keeping agents in the loop.
---
